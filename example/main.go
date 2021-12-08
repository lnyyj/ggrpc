package main

import (
	"log"
	"os"

	"github.com/lnyyj/ggrpc/example/cc"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {
	l := len(os.Args)
	if l < 3 {
		log.Fatalf("输入参数不对")
	}
	mode := os.Args[1] + "_" + os.Args[2]
	addr := "12345"
	if l == 4 {
		addr = os.Args[3]
	}
	cc.Run(mode, addr)
}
