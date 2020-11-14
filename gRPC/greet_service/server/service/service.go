package service

import (
	"context"
	"fmt"
	"github/Ko4s/greet_service/greet"
	"time"
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

//SayManyHello server streaming method
func (s *GreetService) SayManyHello(req *greet.GreetManyRequest, stream greet.GreetService_SayManyHelloServer) error {

	name := req.GetName()
	amount := req.GetAmount() //int32 >- int / int32 / int64

	var i int32
	
	for i = 0; i < amount; i++ {

		res := &greet.GreetManyResponse{
			Result: fmt.Sprintf("Hello %v, i: %v", name, i),
		}

		err := stream.Send(res)

		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
}
