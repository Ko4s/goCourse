package client

import (
	"context"
	"fmt"
	"github/Ko4s/greet_service/greet"
	"io"
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

func (c *Client) SayHelloManyTimes(msg string, times int32) {

	req := &greet.SayHelloManyTimesRequest{
		Name:   msg,
		Amount: times,
	}

	stream, err := c.service.SayHelloManyTimes(context.TODO(), req)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("End of stream data")
			return
		}

		if err != nil {
			return
		}

		fmt.Println(msg.GetResult() + "\n")
	}
}
