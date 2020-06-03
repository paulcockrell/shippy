package main

import (
	"encoding/json"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	userProto "github.com/paulcockrell/shippy/services/user/proto/user"
)

const topic = "user.created"

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.email"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Get the broker instance using our environment variables
	pubsub := service.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe to messages on the broker
	_, err := pubsub.Subscribe(topic, func(e broker.Event) error {
		var user *userProto.User
		if err := json.Unmarshal(e.Message().Body, &user); err != nil {
			return err
		}

		log.Info(user)
		go sendEmail(user)

		return nil
	})

	if err != nil {
		log.Info(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// Dummy emailer
func sendEmail(user *userProto.User) error {
	log.Info("Sending email to: ", user.Name)
	return nil
}
