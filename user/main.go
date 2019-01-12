package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"cinema/user/handler"
	"cinema/user/subscriber"

	example "cinema/user/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.cinema.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.cinema.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("com.cinema.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
