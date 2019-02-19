package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.cinema.srv.movie"),
		micro.Version("latest"),
	)

	if err := cmd.Init(); err != nil {
		log.Fatalf("init cmd error: %v", err)
	}
	// Initialise service
	service.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatalf("run movie service error: %v", err)
	}
}
