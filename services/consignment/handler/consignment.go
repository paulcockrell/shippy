package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	consignment "github.com/paulcockrell/shippy/services/consignment/proto/consignment"

	vesselProto "github.com/paulcockrell/shippy/services/vessel/proto/vessel"
)

type repository interface {
	Create(*consignment.Consignment) (*consignment.Consignment, error)
	GetAll() []*consignment.Consignment
}

// Repository - Dummy in-memory repository
type Repository struct {
	consignments []*consignment.Consignment
}

// Create - Create a consignment and store in memory
func (repo *Repository) Create(consignment *consignment.Consignment) (*consignment.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll - Get all consignments from datastore
func (repo *Repository) GetAll() []*consignment.Consignment {
	return repo.consignments
}

// Consignment - Struct
type Consignment struct {
	Repo         repository
	VesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - Handled by the gRPC server
func (e *Consignment) CreateConsignment(ctx context.Context, req *consignment.Consignment, rsp *consignment.Response) error {
	log.Info("Received Consignment.CreateConsignment request")

	vesselResponse, err := e.VesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Info("Found vessel %s\n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := e.Repo.Create(req)
	if err != nil {
		return err
	}

	rsp.Created = true
	rsp.Consignment = consignment

	return nil
}

// GetConsignments - Get all consignments from datastore
func (e *Consignment) GetConsignments(ctx context.Context, req *consignment.GetRequest, rsp *consignment.Response) error {
	log.Info("Received Consignment.GetConsignments request")
	consignments := e.Repo.GetAll()
	rsp.Consignments = consignments

	return nil
}
