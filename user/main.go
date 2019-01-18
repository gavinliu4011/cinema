package main

import (
	"cinema/user/handler"
	pb "cinema/user/pb/user"
	"cinema/user/wrapper"
	"github.com/getsentry/raven-go"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
)

func init() {
	// 初始化 sentry DSN
	err := raven.SetDSN("https://8d0fb05f0077477baa95cbcb76826867:3b8051df1ead487b94ce053717693b2e@sentry.io/1371816")
	if err != nil {
		log.Fatalf("set sentry DSN error: v%", err)
	}
}

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
	// 必须提前初始化
	err := cmd.Init()
	if err != nil {
		log.Fatalf(" cmd init error: %v", err)
	}
	// Initialise service
	service.Init()

	// Register Handler
	err = pb.RegisterUserServiceHandler(service.Server(), new(handler.UserHandler))
	if err != nil {
		log.Fatalf("register user handler error: %v", err)
	}
	// Register Struct as Subscriber
	//micro.RegisterSubscriber("com.cinema.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("com.cinema.srv.user", service.Server(), subscriber.Handler)

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
