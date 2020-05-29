package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	consignment "consignment/proto/consignment"
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
	repo repository
}

// CreateConsignment - Handled by the gRPC server
func (e *Consignment) CreateConsignment(ctx context.Context, req *consignment.Consignment, rsp *consignment.Response) error {
	log.Info("Received Consignment.CreateConsignment request")

	consignment, err := e.repo.Create(req)
	if err != nil {
		return err
	}

	rsp.Created = true
	rsp.Consignment = consignment

	return nil
}

// GetConsignments - Get all consignments from datastore
func (e *Consignment) GetConsignments(ctx context.Context, req *consignment.GetRequest, rsp *consignment.Response) error {
	consignments := e.repo.GetAll()
	rsp.Consignments = consignments

	return nil
}

/*
// Call is a single request handler called via client.Call or the generated client code
func (e *Consignment) Call(ctx context.Context, req *consignment.Request, rsp *consignment.Response) error {
	log.Info("Received Consignment.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Consignment) Stream(ctx context.Context, req *consignment.StreamingRequest, stream consignment.Consignment_StreamStream) error {
	log.Infof("Received Consignment.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&consignment.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Consignment) PingPong(ctx context.Context, stream consignment.Consignment_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&consignment.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

*/
