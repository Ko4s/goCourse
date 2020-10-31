package service

import (
	"context"
	"github/Ko4s/greet_service/greet"
)

//Service a gRPC service
type GreetService struct {
	greet.UnimplementedGreetServiceServer
}

func NewGreetService() *GreetService {
	return &GreetService{}
}

//SayHello method of rpc service
func (s *GreetService) SayHello(ctx context.Context, request *greet.GreetRequest) (*greet.GreetResponse, error) {
	value := request.GetName()

	return &greet.GreetResponse{
		Greeting: "Hello " + value,
	}, nil
}
