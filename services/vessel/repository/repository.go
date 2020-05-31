package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	vessel "github.com/paulcockrell/shippy/services/vessel/proto/vessel"
)

// Repository -
type Repository interface {
	FindAvailable(ctx context.Context, specification *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
	GetAll(ctx context.Context) ([]*Vessel, error)
}

// MongoRepository -
type MongoRepository struct {
	Collection *mongo.Collection
}

// Vessel - JSON mapping struct
type Vessel struct {
	ID        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerID   string `json:"owner_id"`
}

// MarshalVessel - Converts protobuf to JSON struct
func MarshalVessel(v *vessel.Vessel) *Vessel {
	return &Vessel{
		ID:        v.Id,
		Capacity:  v.Capacity,
		MaxWeight: v.MaxWeight,
		Name:      v.Name,
		Available: v.Available,
		OwnerID:   v.OwnerId,
	}
}

// UnmarshalVessel - Converts JSON struct to protobuf
func UnmarshalVessel(v *Vessel) *vessel.Vessel {
	return &vessel.Vessel{
		Id:        v.ID,
		Capacity:  v.Capacity,
		MaxWeight: v.MaxWeight,
		Name:      v.Name,
		Available: v.Available,
		OwnerId:   v.OwnerID,
	}
}

// UnmarshalVesselCollection -
func UnmarshalVesselCollection(vessels []*Vessel) []*vessel.Vessel {
	collection := make([]*vessel.Vessel, 0)
	for _, vessel := range vessels {
		collection = append(collection, UnmarshalVessel(vessel))
	}

	return collection
}

// Specification - JSON mapping struct
type Specification struct {
	Capacity  int32
	MaxWeight int32
}

// MarshalSpecification - Converts protobuf to JSON struct
func MarshalSpecification(s *vessel.Specification) *Specification {
	return &Specification{
		Capacity:  s.Capacity,
		MaxWeight: s.MaxWeight,
	}
}

// UnmarshalSpecification - Converts JSON struct to protobuf
func UnmarshalSpecification(s *Specification) *vessel.Specification {
	return &vessel.Specification{
		Capacity:  s.Capacity,
		MaxWeight: s.MaxWeight,
	}
}

// FindAvailable -
func (r *MongoRepository) FindAvailable(ctx context.Context, specification *Specification) (*Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			specification.MaxWeight,
		}},
	}}
	vessel := &Vessel{}
	if err := r.Collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}

	return vessel, nil
}

// Create -
func (r *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := r.Collection.InsertOne(ctx, vessel)
	return err
}

// GetAll - Retrieves all vessels (marshaled format)
func (r *MongoRepository) GetAll(ctx context.Context) ([]*Vessel, error) {
	cur, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var vessels []*Vessel
	for cur.Next(ctx) {
		var vessel *Vessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		vessels = append(vessels, vessel)
	}

	return vessels, nil
}
