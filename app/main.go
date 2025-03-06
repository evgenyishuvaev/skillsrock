package main

import (
	"log"
	"skillsrock/internal/config"
	"skillsrock/internal/server"
)

func main() {
	config := config.NewConfig()
	server := server.NewHttpServer(*config)
	log.Fatal(server.Listen(config.App.AsFiberPort()))
}
