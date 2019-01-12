package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"cinema/movie/handler"
	"cinema/movie/subscriber"

	example "cinema/movie/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.cinema.srv.movie"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.cinema.srv.movie", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("com.cinema.srv.movie", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
