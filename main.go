package main

import (
	"github.com/jinho-yoo-jack/gin-tonic-tester/injector"
	"log"
)

func main() {
	server, err := injector.InitializeServer()
	if err != nil {
		log.Fatal("failed to initialize server:", err)
	}
	server.Start(":8080")
}
