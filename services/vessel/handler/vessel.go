package handler

import (
	"context"
	"errors"

	log "github.com/micro/go-micro/v2/logger"
	vessel "github.com/paulcockrell/shippy/services/vessel/proto/vessel"
	repository "github.com/paulcockrell/shippy/services/vessel/repository"
)

// Vessel -
type Vessel struct {
	Repository repository.Repository
}

// FindAvailable - Search vessel repository that matches specification
func (e *Vessel) FindAvailable(ctx context.Context, req *vessel.Specification, rsp *vessel.Response) error {
	log.Info("Received Vessel.FindAvailable request")

	vessel, err := e.Repository.FindAvailable(ctx, repository.MarshalSpecification(req))
	if err != nil {
		return err
	}
	if vessel == nil {
		return errors.New("Could not find vessel")
	}

	rsp.Vessel = repository.UnmarshalVessel(vessel)

	return nil
}

// Create - Create a new vessel model
func (e *Vessel) Create(ctx context.Context, req *vessel.Vessel, rsp *vessel.Response) error {
	log.Info("Received Vessel.Create request")

	if err := e.Repository.Create(ctx, repository.MarshalVessel(req)); err != nil {
		return err
	}

	rsp.Vessel = req

	return nil
}

// GetVessels -
func (e *Vessel) GetVessels(ctx context.Context, req *vessel.GetRequest, rsp *vessel.Response) error {
	log.Info("Received Vessel.GetAll request")

	vessels, err := e.Repository.GetAll(ctx)
	if err != nil {
		return err
	}
	rsp.Vessels = repository.UnmarshalVesselCollection(vessels)

	return nil
}
