package main

import (
	"fmt"
	"github/Ko4s/greet_service/greet"
	"github/Ko4s/greet_service/server/repository"
	"github/Ko4s/greet_service/server/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := "50051"
	lis, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	repo := repository.NewRepository()

	grpcServer := grpc.NewServer(opts...)
	greetService := service.NewGreetService(repo)
	greet.RegisterGreetServiceServer(grpcServer, greetService)

	reflection.Register(grpcServer) //tutaj 

	fmt.Println("Server running ....")
	grpcServer.Serve(lis)
}
