package service

import (
	"context"
	"github/Ko4s/greet_service/greet"
)

type GreetService struct {
	greet.UnimplementedGreetServer
}

func NewGreetService() *GreetService {
	return &GreetService{}
}

func (s *GreetService) SayHello(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {

	name := req.GetName()

	res := &greet.GreetResponse{
		Greeting: "Hello " + name,
	}

	return res, nil
}
