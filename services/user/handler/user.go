package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"
	"golang.org/x/crypto/bcrypt"

	user "github.com/paulcockrell/shippy/services/user/proto/user"
	repository "github.com/paulcockrell/shippy/services/user/repository"
	tokenservice "github.com/paulcockrell/shippy/services/user/tokenservice"
)

// User -
type User struct {
	Repository   repository.Repository
	TokenService *tokenservice.TokenService
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
	pwd := req.Password
	user, err := e.Repository.GetByEmail(req)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)); err != nil {
		return err
	}

	token, err := e.TokenService.Encode(user)
	if err != nil {
		return err
	}

	rsp.Token = token

	return nil
}

// Create -
func (e *User) Create(ctx context.Context, req *user.User, rsp *user.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)

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
