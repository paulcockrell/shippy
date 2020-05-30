package main

import (
	"context"

	"github.com/paulcockrell/shippy/services/consignment/handler"
	"github.com/paulcockrell/shippy/services/consignment/subscriber"

	"os"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	consignment "github.com/paulcockrell/shippy/services/consignment/proto/consignment"

	vesselProto "github.com/paulcockrell/shippy/services/vessel/proto/vessel"

	repository "github.com/paulcockrell/shippy/services/consignment/repository"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
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

	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.consignment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//consignment.RegisterShippingServiceHandler(service.Server(), new(handler.Consignment{repo}))
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
