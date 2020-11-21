package main

import (
	"fmt"
	"github/Ko4s/greet_service/client/client"
)

func main() {

	address := "localhost:50051"
	c, err := client.NewClient(address)

	if err != nil {
		panic(err)
	}

	names := []string{"Piotrek", "Olga", "Aneta", "Szymon"}

	r, err := c.GreetManyUsers(names)

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
