package main

import (
	"encoding/json"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/client/selector/static/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	userProto "github.com/paulcockrell/shippy/services/user/proto/user"
)

const topic = "user.created"

func main() {
	/*** Setup Service ***/

	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// New Service
	service := micro.NewService(
		micro.Registry(r),
		micro.Selector(s),
		micro.Name("com.foo.service.email"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.AfterStart(func() error {
		brk := service.Options().Broker
		if err := brk.Connect(); err != nil {
			log.Fatalf("Broker connect error: %v", err)
		}

		go sub(brk)

		return nil
	}))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func sub(brk broker.Broker) {
	// Subscribe to messages on the broker
	_, err := brk.Subscribe(topic, func(p broker.Event) error {
		log.Info("Got an event")

		var user *userProto.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}

		log.Info(user)

		go sendEmail(user)

		return nil
	}, broker.Queue(topic))

	if err != nil {
		log.Info(err)
	}
}

// Dummy emailer
func sendEmail(user *userProto.User) error {
	log.Info("Sending email to: ", user.Name)
	return nil
}
