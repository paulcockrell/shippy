package main

import (
	"github.com/paulcockrell/shippy/services/consignment/handler"
	"github.com/paulcockrell/shippy/services/consignment/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	consignment "github.com/paulcockrell/shippy/services/consignment/proto/consignment"

	vesselProto "github.com/paulcockrell/shippy/services/vessel/proto/vessel"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.consignment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	repo := &handler.Repository{}

	// Register Handler
	//consignment.RegisterShippingServiceHandler(service.Server(), new(handler.Consignment{repo}))
	consignment.RegisterShippingServiceHandler(service.Server(), &handler.Consignment{Repo: repo, VesselClient: vesselProto.VesselServiceClient})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.service.consignment", service.Server(), new(subscriber.Consignment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
