package main

import (
	"consignment/handler"
	"consignment/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	consignment "consignment/proto/consignment"
)

type repository interface {
	Create(consignment.Consignment) (*consignment.Consignment, error)
	GetAll() []*consignment.Consignment
}

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
	consignment.RegisterShippingServiceHandler(service.Server(), &handler.Consignment{Repo: repo})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.service.consignment", service.Server(), new(subscriber.Consignment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
