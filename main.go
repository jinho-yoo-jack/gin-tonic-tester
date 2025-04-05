package main

import (
	c "ginTonicProject/config"
	s "ginTonicProject/server"
)

func main() {
	config, err := c.LoadConfig()
	if err != nil {
		panic(err)
	}
	server, err := s.NewServer(config, c.DB)
	err = server.Start(":8080")
	if err != nil {
		return
	}
}
