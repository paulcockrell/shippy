package handler

import (
	"context"
	"errors"

	log "github.com/micro/go-micro/v2/logger"

	consignment "github.com/paulcockrell/shippy/services/consignment/proto/consignment"

	vesselProto "github.com/paulcockrell/shippy/services/vessel/proto/vessel"

	repository "github.com/paulcockrell/shippy/services/consignment/repository"
)

// Consignment - Struct
type Consignment struct {
	Repository   repository.Repository
	VesselClient vesselProto.VesselService
}

// CreateConsignment - Handled by the gRPC server
func (e *Consignment) CreateConsignment(ctx context.Context, req *consignment.Consignment, rsp *consignment.Response) error {
	log.Info("Received Consignment.CreateConsignment request")

	vesselResponse, err := e.VesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if vesselResponse == nil {
		return errors.New("error fetching vessel, returned nil")
	}
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err := e.Repository.Create(ctx, repository.MarshalConsignment(req)); err != nil {
		return err
	}

	rsp.Created = true
	rsp.Consignment = req

	return nil
}

// GetConsignments - Get all consignments from datastore
func (e *Consignment) GetConsignments(ctx context.Context, req *consignment.GetRequest, rsp *consignment.Response) error {
	log.Info("Received Consignment.GetConsignments request")

	consignments, err := e.Repository.GetAll(ctx)
	if err != nil {
		return err
	}
	rsp.Consignments = repository.UnmarshalConsignmentCollection(consignments)

	return nil
}
