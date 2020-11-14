package main

import "github/Ko4s/calculator_service/client/client"

func main() {

	address := "localhost:50051"
	c, err := client.NewClient(address)

	if err != nil {
		panic(err)
	}

	c.Sum(1, 2)
}
