package client

import (
	"context"
	"github/Ko4s/greet_service/greet"
	"log"

	"google.golang.org/grpc"
)

type Client struct {
	service greet.GreetServiceClient
}

func NewClient(address string) *Client {

	//cc -> Client connection
	cc, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	greetClient := greet.NewGreetServiceClient(cc)

	return &Client{
		service: greetClient,
	}

}

func (c *Client) SayHello(msg string) string {

	req := greet.GreetRequest{
		Name: msg,
	}

	res, err := c.service.SayHello(context.TODO(), &req)

	if err != nil {
		log.Fatalln(err)
	}

	return res.GetGreeting()
}
