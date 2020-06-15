package main

import (
	"context"
	"errors"

	"github.com/paulcockrell/shippy/services/consignment/handler"
	"github.com/paulcockrell/shippy/services/consignment/subscriber"

	"os"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"

	consignment "github.com/paulcockrell/shippy/services/consignment/proto/consignment"

	userProto "github.com/paulcockrell/shippy/services/user/proto/user"
	vesselProto "github.com/paulcockrell/shippy/services/vessel/proto/vessel"

	"github.com/micro/go-plugins/client/selector/static/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	repository "github.com/paulcockrell/shippy/services/consignment/repository"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	log.Info("Consignment service booting")
	/*** Setup DB ***/

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")
	repo := &repository.MongoRepository{Collection: consignmentCollection}

	/*** Setup Service ***/

	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// New Service
	service := micro.NewService(
		micro.Registry(r),
		micro.Selector(s),
		micro.Name("com.foo.service.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	// Initialise service
	service.Init()

	// Register Handler
	vesselClient := vesselProto.NewVesselService("com.foo.service.vessel", service.Client())
	h := &handler.Consignment{Repository: repo, VesselClient: vesselClient}
	consignment.RegisterShippingServiceHandler(service.Server(), h)

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.service.consignment", service.Server(), new(subscriber.Consignment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// AuthWrapper - Validate token with user service before handling request
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]

		authClient := userProto.NewUserService("com.foo.service.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userProto.Token{
			Token: token,
		})
		if err != nil {
			return err
		}

		err = fn(ctx, req, rsp)

		return err
	}
}
