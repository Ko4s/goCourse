package main

import "github/Ko4s/greet_service/client/client"

func main() {

	address := "localhost:50051"
	c, err := client.NewClient(address)

	if err != nil {
		panic(err)
	}

	c.SayHello("Piotrek")
	c.SayHello("Szymon")
	c.SayHello("Aneta")
}
