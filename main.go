package main

import (
	"log"
	"mark99/injector"
)

func main() {
	server, err := injector.InitializeServer()
	if err != nil {
		log.Fatal("failed to initialize server:", err)
	}
	server.Start(":8080")
}
