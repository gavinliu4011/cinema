package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"cinema/ticketing/handler"
	"cinema/ticketing/subscriber"

	example "cinema/ticketing/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.cinema.srv.ticketing"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.cinema.srv.ticketing", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("com.cinema.srv.ticketing", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
