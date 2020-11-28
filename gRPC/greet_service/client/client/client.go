package client

import (
	"context"
	"fmt"
	"github/Ko4s/greet_service/greet"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (c *Client) GreetManyTimes(names []string) ([]string, error) {

	clientStream, err := c.service.GreetManyTimes(context.TODO())

	if err != nil {
		return nil, err
	}

	//1. Potrzebujemy gorutyn
	//zrobienia chanela
	waitc := make(chan string)

	go func() {
		for _, name := range names {

			req := greet.GreetRequest{
				Name: name,
			}

			clientStream.Send(&req)
			time.Sleep(time.Second)
		}
		clientStream.CloseSend()
	}()

	//funkcja do czytania odpiedzi z serwera

	go func() {
		for {
			res, err := clientStream.Recv()

			if err == io.EOF {
				fmt.Println(err)
				break
			}

			if err != nil {
				log.Fatalf("%v", err)
			}
			waitc <- res.GetGreeting()

		}
		close(waitc)
	}()

	var l = []string{}

	for el := range waitc {
		l = append(l, el)
	}

	return l, nil
}

func (c *Client) MatchNameWithData(name string, t time.Duration) {

	req := &greet.MatchNameWithDataRequest{
		Name: name,
	}

	ctxWithTimout, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	res, err := c.service.MatchNameWithData(ctxWithTimout, req)

	errStatus, _ := status.FromError(err)

	if errStatus.Code() == codes.NotFound {
		fmt.Println("Nie udało się coś znaleść resourca")
	}

	if errStatus.Code() == codes.Canceled {
		fmt.Println("Serwer nie ogarnał  w tri miga")
	}

	if err != nil {
		log.Println(err)
		log.Fatalln(err.Error())
	}

	fmt.Println(res.GetAge())
}
