package main

import (
	"log"
	"os"
)

func main() {
	mode := "simple"
	if len(os.Args) == 2 {
		mode = os.Args[1]
	}
	addr := ":12345"

	log.Printf("example grpc server %+v starting...\r\n", addr)

	switch mode {
	case "simple":
		simpleServer(addr)
	}
}
