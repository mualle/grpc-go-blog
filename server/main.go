package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {

	//if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("usage: server [IP_ADDR]")
	}

	addr := args[0]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer func(lis net.Listener) {
		if err := lis.Close(); err != nil {
			log.Fatalf("Unexpected error: %v\n", err)
		}
	}(lis)

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	//Registration of endpoints
	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
