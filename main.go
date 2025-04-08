package main

import (
	"github.com/jinho-yoo-jack/gin-tonic-tester/internal/di"
	"log"
)

func main() {
	server, err := di.InitializeServer()
	if err != nil {
		log.Fatal("failed to initialize server:", err)
	}
	server.Start(":8080")
}
