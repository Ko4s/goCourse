package service

import (
	"context"
	"github/Ko4s/greet_service/greet"
	"io"
	"strconv"
	"strings"
	"time"
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

func (s *GreetService) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.Greet_GreetManyTimesServer) error {

	name := req.GetName()
	times := int(req.GetTimes())

	for i := 0; i < times; i++ {
		res := greet.GreetManyTimesResponse{
			Msg: "Hello " + name + " " + strconv.Itoa(i+1),
		}

		err := stream.Send(&res)

		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
}
