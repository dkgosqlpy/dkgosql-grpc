package main

import (
	"fmt"
	"log"
	"net"
	"github.com/dkgosqlpy/dkgosql-grpc/greetpb"
)

type server struct{}

func main() {
	fmt.Println("Hello World!")
	listner, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grcp.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listner); err !nil{
		log.Fatalf("Failed to serve: %v", err)
	}

}
