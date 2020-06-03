package handler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"golang.org/x/crypto/bcrypt"

	user "github.com/paulcockrell/shippy/services/user/proto/user"
	repository "github.com/paulcockrell/shippy/services/user/repository"
	tokenservice "github.com/paulcockrell/shippy/services/user/tokenservice"
)

const topic = "user.created"

// User -
type User struct {
	Repository   repository.Repository
	TokenService *tokenservice.TokenService
	PubSub       broker.Broker
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
	if err := e.publishEvent(req); err != nil {
		return err
	}

	return nil
}

func (e *User) publishEvent(user *user.User) error {
	log.Info("Publishing event!")

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Create a broker message
	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	if err := e.PubSub.Publish(topic, msg); err != nil {
		log.Infof("[pub] failed: %v", err)
		return err
	}

	return nil
}

// ValidateToken -
func (e *User) ValidateToken(ctx context.Context, req *user.Token, rsp *user.Token) error {
	// Decode token
	claims, err := e.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	rsp.Valid = true

	return nil
}
