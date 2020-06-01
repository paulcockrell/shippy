package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/paulcockrell/shippy/services/user/handler"
	user "github.com/paulcockrell/shippy/services/user/proto/user"
	repository "github.com/paulcockrell/shippy/services/user/repository"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	defer db.Close()

	db.AutoMigrate(&user.User{})
	repository := &repository.UserRepository{db}

	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserServiceHandler(service.Server(), &handler.User{repository})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
