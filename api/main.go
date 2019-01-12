package main

import (
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"cinema/api/handler"
	"cinema/api/client"

	example "cinema/api/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.cinema.api.api"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.ExampleWrapper(service)),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
