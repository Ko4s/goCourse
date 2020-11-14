package client

import (
	"context"
	"fmt"
	"github/Ko4s/calculator_service/calc"
	"log"

	"google.golang.org/grpc"
)

type Client struct {
	service calc.CalcClient
}

func NewClient(address string) (*Client, error) {

	cc, err := grpc.Dial(address, grpc.WithInsecure()) //łączymy się z serwerem gRPC po danym adresie, w opjach wybraliśmy bez zabepieczeń czyli nasz ruch jest nie szyforwany

	if err != nil {
		return nil, err
	}

	sumClient := calc.NewCalcClient(cc) //tutaj tworzymy nasz "serwis kliencki"

	return &Client{
		service: sumClient,
	}, nil
}

func (c *Client) Sum(num1, num2 int32) {

	req := &calc.SumRequest{
		Number1: num1,
		Number2: num2,
	}
	ctx := context.TODO()

	res, err := c.service.Sum(ctx, req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetResult())
}
