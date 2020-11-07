package main

import "github/Ko4s/greet_service/client/client"

func main() {

	address := "localhost:50051"

	c := client.NewClient(address)

	c.SayManyHello("Piotrek", 17)
}
