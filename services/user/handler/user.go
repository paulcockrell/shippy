package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	user "github.com/paulcockrell/shippy/services/user/proto/user"
	repository "github.com/paulcockrell/shippy/services/user/repository"
)

// User -
type User struct {
	Repository repository.Repository
}

// GetAll -
func (e *User) GetAll(ctx context.Context, req *user.Request, rsp *user.Response) error {
	users, err := e.Repository.GetAll(ctx)
	if err != nil {
		return err
	}

	rsp.Users = users

	return nil
}

// Get -
func (e *User) Get(ctx context.Context, req *user.User, rsp *user.Response) error {
	log.Info("Received User.Get request")

	user, err := e.Repository.Get(ctx, req.Id)
	if err != nil {
		return err
	}
	rsp.User = user
	return nil
}

// Auth -
func (e *User) Auth(ctx context.Context, req *user.User, rsp *user.Token) error {
	_ /*user*/, err := e.Repository.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	rsp.Token = "testingabc"

	return nil
}

// Create -
func (e *User) Create(ctx context.Context, req *user.User, rsp *user.Response) error {
	if err := e.Repository.Create(req); err != nil {
		return err
	}

	rsp.User = req

	return nil
}

// ValidateToken -
func (e *User) ValidateToken(ctx context.Context, req *user.Token, rsp *user.Token) error {
	return nil
}
