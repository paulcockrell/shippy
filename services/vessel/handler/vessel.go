package handler

import (
	"context"
	"errors"

	vessel "github.com/paulcockrell/shippy/services/vessel/proto/vessel"
)

type repository interface {
	FindAvailable(*vessel.Specification) (*vessel.Vessel, error)
}

// Repository - Dummy repository for vessels
type Repository struct {
	Vessels []*vessel.Vessel
}

// Vessel -
type Vessel struct {
	Repo repository
}

// FindAvailable - Return a vessel from the repository that meets the specification
func (repo *Repository) FindAvailable(spec *vessel.Specification) (*vessel.Vessel, error) {
	for _, vessel := range repo.Vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// FindAvailable - Search vessel repository that matches specification
func (e *Vessel) FindAvailable(ctx context.Context, req *vessel.Specification, rsp *vessel.Response) error {
	vessel, err := e.Repo.FindAvailable(req)
	if err != nil {
		return err
	}

	rsp.Vessel = vessel

	return nil
}
