package service

import (
	"context"
	"github/Ko4s/greet_service/greet"
	"io"
	"strings"
)

//GreetService implementation of GreetServer grpc interface
type GreetService struct {
	greet.UnimplementedGreetServer
}

//NewGreetService constructor of GreetService
func NewGreetService() *GreetService {
	return &GreetService{}
}

//SayHello says hello to a user
func (s *GreetService) SayHello(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {

	name := req.GetName()

	res := &greet.GreetResponse{
		Greeting: "Hello " + name,
	}

	return res, nil
}

//GreetManyUsers accpet stream of names and returns greeting
func (s *GreetService) GreetManyUsers(stream greet.Greet_GreetManyUsersServer) error {

	names := []string{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		names = append(names, req.GetName())
	}

	res := greet.GreetManyUsersResponse{
		Result: "Hello " + strings.Join(names, ", "),
	}

	return stream.SendAndClose(&res)
}
