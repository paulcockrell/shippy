package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	consignment "consignment/proto/consignment"
)

type Consignment struct{}

func (e *Consignment) Handle(ctx context.Context, msg *consignment.Consignment) error {
	log.Info("Handler Received message: ", msg.GetId)
	return nil
}

func Handler(ctx context.Context, msg *consignment.Consignment) error {
	log.Info("Function Received message: ", msg.GetId)
	return nil
}
