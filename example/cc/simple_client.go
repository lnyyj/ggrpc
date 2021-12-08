package cc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lnyyj/ggrpc/example/pb"
	"google.golang.org/grpc/resolver"

	_ "github.com/yangxikun/go-grpc-client-side-lb-example/resolver/dns"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

const (
	defaultName = "rokety"
)

func SimpleClient(address string) {
	// Set resolver
	resolver.SetDefaultScheme("custom_dns")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithBlock(), grpc.WithBackoffMaxDelay(time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	for range time.Tick(time.Second) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
		if err != nil {
			log.Printf("could not greet: %v\n", err)
		} else {
			log.Printf("Greeting: %s", r.Message)
		}
		cancel()
	}
}
