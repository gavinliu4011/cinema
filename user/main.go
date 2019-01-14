package main

import (
	"cinema/user/handler"
	pb "cinema/user/pb/user"
	"cinema/user/wrapper"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.cinema.srv.user"),
		micro.Version("latest"),
		micro.WrapHandler(wrapper.HandlerWrapper),
		micro.WrapClient(wrapper.LogClientWrap),
		micro.WrapCall(wrapper.CallFuncWrap),
		//micro.WrapHandler(ratelimit.NewHandlerWrapper(&rate.Bucket{}, false)),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := pb.RegisterUserServiceHandler(service.Server(), new(handler.UserHandler))
	if err != nil {
		log.Fatalf("register user handler error: %v", err)
	}
	// Register Struct as Subscriber
	//micro.RegisterSubscriber("com.cinema.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("com.cinema.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
