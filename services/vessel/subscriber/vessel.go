package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	vessel "vessel/proto/vessel"
)

type Vessel struct{}

func (e *Vessel) Handle(ctx context.Context, msg *vessel.Specification) error {
	log.Info("Handler Received message: ", msg.GetCapacity)
	return nil
}

func Handler(ctx context.Context, msg *vessel.Specification) error {
	log.Info("Function Received message: ", msg.GetCapacity)
	return nil
}
