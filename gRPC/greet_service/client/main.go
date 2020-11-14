package main

import (
	"fmt"
	"github/Ko4s/greet_service/client/client"
)

func main() {

	address := "localhost:50051"

	c := client.NewClient(address)

	names := []string{"Piotrek", "Szymon", "Aneta", "Olga"}
	s, _ := c.GreetManyUsers(names)
	fmt.Println(s)
}
