package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

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
	return &pb.HelloReply{Message: "Hello " + in.Name + "! From " + GetIP()}, nil
}

func simple_server() {
	// simulate busy server
	stuckDuration = time.Duration(rand.NewSource(time.Now().UnixNano()).Int63()%2) * time.Second

	if stuckDuration == time.Second {
		log.Println("I will stuck one second!!!")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GetIP() string {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}
			if ip.String() != "127.0.0.1" {
				return ip.String()
			}
		}
	}
	return ""
}
