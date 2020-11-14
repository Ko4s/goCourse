package service

import (
	"context"
	"fmt"
	"github/Ko4s/greet_service/greet"
	"io"
	"strings"
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

func (s *GreetService) GreetManyUsers(stream greet.GreetService_GreetManyUsersServer) error {

	var names = []string{}
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

	res := &greet.GreetResponse{
		Greeting: fmt.Sprintf("Hello %v ", strings.Join(names, " ")),
	}

	fmt.Println(fmt.Sprintf("Hello %v ", strings.Join(names, " ")))
	if err := stream.SendAndClose(res); err != nil {
		return err
	}

	return nil
}
