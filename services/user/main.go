package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/paulcockrell/shippy/services/user/handler"
	user "github.com/paulcockrell/shippy/services/user/proto/user"
	repository "github.com/paulcockrell/shippy/services/user/repository"
	tokenservice "github.com/paulcockrell/shippy/services/user/tokenservice"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	defer db.Close()

	db.AutoMigrate(&user.User{})
	repo := &repository.UserRepository{db}
	ts := &tokenservice.TokenService{repo}

	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.user"),
		micro.Version("latest"),
	)

	// var pubsub broker.Broker

	// Initialise service
	// service.Init(micro.AfterStart(func() error {
	// 	pubsub = service.Options().Broker
	// 	if err := pubsub.Connect(); err != nil {
	// 		log.Fatalf("Broker connect error: %v", err)
	// 	}

	// 	return nil
	// }))

	service.Init()

	// Get an instance of the broker using our defaults
	pubsub := service.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatalf("Broker not connected error: %v", err)
	}

	// Register Handler
	h := &handler.User{
		Repository:   repo,
		TokenService: ts,
		PubSub:       pubsub,
	}
	user.RegisterUserServiceHandler(service.Server(), h)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
