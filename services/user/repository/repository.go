package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	log "github.com/micro/go-micro/v2/logger"
	user "github.com/paulcockrell/shippy/services/user/proto/user"
)

// Repository - Interface
type Repository interface {
	GetAll(ctx context.Context) ([]*user.User, error)
	Get(ctx context.Context, id string) (*user.User, error)
	GetByEmailAndPassword(user *user.User) (*user.User, error)
	Create(user *user.User) error
}

// UserRepository - Holds mongo collection
type UserRepository struct {
	Db *gorm.DB
}

// GetAll -
func (r *UserRepository) GetAll(ctx context.Context) ([]*user.User, error) {
	var users []*user.User
	if err := r.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Get -
func (r *UserRepository) Get(ctx context.Context, id string) (*user.User, error) {
	var user *user.User
	user.Id = id
	if err := r.Db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmailAndPassword -
func (r *UserRepository) GetByEmailAndPassword(user *user.User) (*user.User, error) {
	if err := r.Db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Create -
func (r *UserRepository) Create(user *user.User) error {
	log.Info("Create user", user)
	err := r.Db.Create(user).Error
	return err
}
