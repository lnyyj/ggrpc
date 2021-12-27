package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/lnyyj/ggrpc/example/lib"
	"github.com/lnyyj/ggrpc/example/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var stuckDuration time.Duration

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	time.Sleep(stuckDuration)
	return &pb.HelloReply{Message: "Hello " + in.Name + "! From " + lib.GetIP()}, nil
}

func simpleServer(addr string) {
	// simulate busy server
	stuckDuration = time.Duration(rand.NewSource(time.Now().UnixNano()).Int63()%2) * time.Second
	if stuckDuration == time.Second {
		log.Println("I will stuck one second!!!")
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
