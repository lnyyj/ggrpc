package cc

import (
	"log"
	"strings"
)

func Run(mode, port string) {
	mode = strings.ToLower(mode)
	log.Println("cc run mode "+mode, " port: ", port)
	switch mode {
	case "simple_client":
		SimpleClient("localhost:" + port)
	case "simple_server":
		SimpleServer(":" + port)
	}
}
