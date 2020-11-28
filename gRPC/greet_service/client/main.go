package main

import (
	"github/Ko4s/greet_service/client/client"
	"time"
)

func main() {

	address := "localhost:50051"

	c := client.NewClient(address)

	c.MatchNameWithData("Piotrek", time.Second*7)
}
