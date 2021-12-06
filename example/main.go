package main

import (
	"flag"
	"log"
)

var (
	port = ":50051"

	address string
	timeout int

	defaultName = "1234"
)

func init() {
	flag.IntVar(&timeout, "timeout", 1, "greet rpc call timeout")
	flag.StringVar(&address, "address", "localhost:50051", "grpc server addr")
	flag.Parse()

	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {

}
