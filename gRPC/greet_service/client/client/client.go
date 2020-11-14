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

func (c *Client) SayManyHello(name string, amount int32) {
	req := &greet.GreetManyRequest{
		Name:   name,
		Amount: amount,
	}

	resStream, err := c.service.SayManyHello(context.TODO(), req)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := resStream.Recv()

		if err == io.EOF {
			//ten error dostaniemy gdy streamowanie się zakonczy
			fmt.Println("Data streaming eneded...")
			return
		}

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(res.GetResult())
	}
}

func (c *Client) GreetManyUsers(names []string) (string, error) {

	inputStream, err := c.service.GreetManyUsers(context.TODO())

	if err != nil {
		return "", err
	}

	for _, name := range names {
		req := &greet.GreetRequest{
			Name: name,
		}

		inputStream.Send(req)
	}

	res, err := inputStream.CloseAndRecv()

	if err != nil {
		return "", err
	}

	return res.GetGreeting(), nil
}
