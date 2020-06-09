package main

import (
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	proto "github.com/paulcockrell/shippy/services/user-api/proto/user"

	"context"
)

// Example - empty struct
type Example struct{}

// Foo - empty struct
type Foo struct{}

// Call - Call method on the example struct
func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("Received Example.Call request")

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}

	rsp.Message = "got your request " + req.Name

	return nil
}

// Bar - bar black sheep
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("Received Foo.Bar request")

	// noop

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("fart.bob.pants.example"),
	)

	service.Init()

	proto.RegisterExampleHandler(service.Server(), new(Example))
	proto.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run; err != nil {
		log.Print(err())
		os.Exit(1)
	}
}
