package repository

import (
	"context"

	consignment "github.com/paulcockrell/shippy/services/consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

// Consignment - JSON mapping struct
type Consignment struct {
	ID          string     `json:"id"`
	Weight      int32      `json:"weight"`
	Description string     `json:"description"`
	Containers  Containers `json:"containers"`
	VesselID    string     `json:"vessel_id"`
}

// Container - JSON mapping struct
type Container struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	UserID     string `json:"user_id"`
}

// Containers - Array of JSON mapping structs
type Containers []*Container

// MarshalContainerCollection - Converts protobuf to JSON struct
func MarshalContainerCollection(containers []*consignment.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}

	return collection
}

// UnmarshalContainerCollection - Converts JSON struct to protobuf
func UnmarshalContainerCollection(containers []*Container) []*consignment.Container {
	collection := make([]*consignment.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}

	return collection
}

// UnmarshalContainer - Converts JSON struct to protobuf
func UnmarshalContainer(c *Container) *consignment.Container {
	return &consignment.Container{
		Id:         c.ID,
		CustomerId: c.CustomerID,
		UserId:     c.UserID,
	}
}

// MarshalContainer - Converts protobuf to JSON struct
func MarshalContainer(c *consignment.Container) *Container {
	return &Container{
		ID:         c.Id,
		CustomerID: c.CustomerId,
		UserID:     c.UserId,
	}
}

// MarshalConsignment - Converts protobuf to JSON struct
func MarshalConsignment(c *consignment.Consignment) *Consignment {
	return &Consignment{
		ID:          c.Id,
		Weight:      c.Weight,
		Description: c.Description,
		Containers:  MarshalContainerCollection(c.Containers),
		VesselID:    c.VesselId,
	}
}

// UnmarshalConsignment - Converts JSON struct to protobuf
func UnmarshalConsignment(c *Consignment) *consignment.Consignment {
	return &consignment.Consignment{
		Id:          c.ID,
		Weight:      c.Weight,
		Description: c.Description,
		Containers:  UnmarshalContainerCollection(c.Containers),
		VesselId:    c.VesselID,
	}
}

// UnmarshalConsignmentCollection - Converts JSON struct to protobuf
func UnmarshalConsignmentCollection(consignments []*Consignment) []*consignment.Consignment {
	collection := make([]*consignment.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
	}

	return collection
}

// Repository - Interface
type Repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

// MongoRepository - Holds mongo collection
type MongoRepository struct {
	Collection *mongo.Collection
}

// Create - Inserts marshaled consignment
func (r *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := r.Collection.InsertOne(ctx, consignment)
	return err
}

// GetAll - Retrieves all consignments (marshaled format)
func (r *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := r.Collection.Find(ctx, nil, nil)
	var consignments []*Consignment
	for cur.Next(ctx) {
		var consignment *Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}

	return consignments, err
}
