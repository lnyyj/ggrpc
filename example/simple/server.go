package simple

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

func Client(address string) {
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
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
		if err != nil {
			log.Printf("could not greet: %v\n", err)
		} else {
			log.Printf("Greeting: %s", r.Message)
		}
		cancel()
	}
}
