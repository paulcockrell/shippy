package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-plugins/client/selector/static/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	"github.com/paulcockrell/shippy/services/user/handler"
	user "github.com/paulcockrell/shippy/services/user/proto/user"
	repository "github.com/paulcockrell/shippy/services/user/repository"
	tokenservice "github.com/paulcockrell/shippy/services/user/tokenservice"
)

func main() {
	/*** Setup DB ***/

	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	defer db.Close()

	db.AutoMigrate(&user.User{})
	repo := &repository.UserRepository{db}
	ts := &tokenservice.TokenService{repo}

	/*** Setup service ***/

	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// New Service
	service := micro.NewService(
		micro.Registry(r),
		micro.Selector(s),
		micro.Name("com.foo.service.user"),
		micro.Version("latest"),
	)

	service.Init()

	// Register Handler
	publisher := micro.NewEvent("user.created", service.Client())
	h := &handler.User{
		Repository:   repo,
		TokenService: ts,
		Publisher:    publisher,
	}
	user.RegisterUserServiceHandler(service.Server(), h)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
