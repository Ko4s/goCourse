package main

import (
	"fmt"
	"github/Ko4s/calculator_service/calc"
	"github/Ko4s/calculator_service/server/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port string
var address string

func init() {
	port = "50051"
	address = "localhost:" + port
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption //tu tablica/slice jest pusta

	grpcServer := grpc.NewServer(opts...)   //zrobiłem sobie server gRPC
	calcService := service.NewCalcService() //zrobiłem sobie moją implementację servicu, czyli tego structa GreetService z metoda SayHello

	calc.RegisterCalcServer(grpcServer, calcService)

	fmt.Println("Server up and running...")
	grpcServer.Serve(lis)
}
