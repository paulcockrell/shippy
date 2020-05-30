package main

import (
	"github.com/paulcockrell/shippy/services/vessel/handler"
	"github.com/paulcockrell/shippy/services/vessel/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	vessel "github.com/paulcockrell/shippy/services/vessel/proto/vessel"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.vessel"),
		micro.Version("latest"),
	)

	// Lets add some data to the vessel repo
	vessels := []*vessel.Vessel{
		{
			Id: "vessel001", Name: "Boaty McBoatFace",
			MaxWeight: 200000, Capacity: 500,
		},
	}
	repo := &handler.Repository{Vessels: vessels}

	// Initialise service
	service.Init()

	// Register Handler
	vessel.RegisterVesselServiceHandler(service.Server(), &handler.Vessel{Repo: repo})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.service.vessel", service.Server(), new(subscriber.Vessel))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
