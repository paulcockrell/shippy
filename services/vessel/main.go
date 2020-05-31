package main

import (
	"context"
	"os"

	"github.com/paulcockrell/shippy/services/vessel/handler"
	"github.com/paulcockrell/shippy/services/vessel/repository"
	"github.com/paulcockrell/shippy/services/vessel/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	vessel "github.com/paulcockrell/shippy/services/vessel/proto/vessel"
)

const (
	defaultHost = "mongodb://localhost:27017"
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

	vesselCollection := client.Database("shippy").Collection("vessels")
	repo := &repository.MongoRepository{Collection: vesselCollection}

	/*** Setup Service ***/

	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.vessel"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	h := &handler.Vessel{Repository: repo}
	vessel.RegisterVesselServiceHandler(service.Server(), h)

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.service.vessel", service.Server(), new(subscriber.Vessel))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
