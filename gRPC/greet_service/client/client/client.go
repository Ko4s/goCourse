package client

import (
	"context"
	"fmt"
	"github/Ko4s/greet_service/greet"
	"log"

	"google.golang.org/grpc"
)

type Client struct {
	service greet.GreetClient
}

func NewClient(address string) (*Client, error) {

	cc, err := grpc.Dial(address, grpc.WithInsecure()) //łączymy się z serwerem gRPC po danym adresie, w opjach wybraliśmy bez zabepieczeń czyli nasz ruch jest nie szyforwany

	if err != nil {
		return nil, err
	}

	greetClient := greet.NewGreetClient(cc) //tutaj tworzymy nasz "serwis kliencki"

	return &Client{
		service: greetClient,
	}, nil
}

func (c *Client) SayHello(name string) {

	req := &greet.GreetRequest{
		Name: name,
	}
	ctx := context.TODO()

	res, err := c.service.SayHello(ctx, req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetGreeting())
}

func (c *Client) GreetManyUsers(names []string) (string, error) {

	stream, err := c.service.GreetManyUsers(context.TODO())

	if err != nil {
		return "", err
	}

	for _, name := range names {
		req := greet.GreetManyUsersRequest{
			Name: name,
		}

		stream.Send(&req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		return "", err
	}

	return res.GetResult(), nil
}
