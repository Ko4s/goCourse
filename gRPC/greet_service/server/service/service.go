package service

import (
	"context"
	"fmt"
	"github/Ko4s/greet_service/greet"
	"github/Ko4s/greet_service/server/repository"
	"io"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Service a gRPC service
type GreetService struct {
	repo repository.Repository
	greet.UnimplementedGreetServiceServer
}

func NewGreetService(repo repository.Repository) *GreetService {
	return &GreetService{repo: repo}
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

//GreetManyUsers client streaming method
func (s *GreetService) GreetManyUsers(stream greet.GreetService_GreetManyUsersServer) error {

	names := []string{}

	for {
		streamReq, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		// if streamReq.GetName() == "Piotrek" {
		// 	res := &greet.GreetResponse{
		// 		Greeting: fmt.Sprintf("Die %v", streamReq.GetName()),
		// 	}
		// 	return stream.SendAndClose(res)
		// }

		names = append(names, streamReq.GetName())
	}

	res := &greet.GreetResponse{
		Greeting: fmt.Sprintf("Hello %v", strings.Join(names, ", ")),
	}

	return stream.SendAndClose(res)
}

//GreetManyTimes streamming service
func (s *GreetService) GreetManyTimes(serverStream greet.GreetService_GreetManyTimesServer) error {

	for {
		req, err := serverStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		name := req.GetName()

		res := greet.GreetResponse{
			Greeting: "Hello " + name,
		}

		serverStream.Send(&res)
	}

	return nil
}

//MatchNameWithData Example of error statuses in gRPC
func (s *GreetService) MatchNameWithData(ctx context.Context, req *greet.MatchNameWithDataRequest) (*greet.MatchNameWithDataResponse, error) {

	name := req.GetName()
	//timeout ktorym będę blokował funckje

	for i := 0; i < 10; i++ {

		if ctx.Err() == context.Canceled {
			return nil, status.Error(codes.Canceled, ":(")
		}

		fmt.Println(i)
		time.Sleep(time.Second)
	}

	res := &greet.MatchNameWithDataResponse{
		Age:     17,
		PetName: name,
	}

	return res, nil
}
