package service

import (
	"context"
	"fmt"
	"github/Ko4s/greet_service/greet"
	"time"
)

//GreetService a gRPC service
type GreetService struct {
	greet.UnimplementedGreetServiceServer
}

//NewGreetService GreetService Constructor
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

//SayHelloManyTimes is Streaming method
func (s *GreetService) SayHelloManyTimes(req *greet.SayHelloManyTimesRequest, stream greet.GreetService_SayHelloManyTimesServer) error {

	//1 Pobieramy responsa od klienta będzie jeden
	reqAmount := req.GetAmount()
	reqName := req.GetName()

	var i int32
	for i = 0; i < reqAmount; i++ {

		res := &greet.SayHelloManyTimesResponse{
			Result: fmt.Sprintf("Hello %v times: %v", reqName, i),
		}

		err := stream.Send(res)

		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
}
