package main

import (
	"cinema/api/router"
	"github.com/micro/go-web"
	"log"
	"net/http"
	"time"
)

func main() {
	// 创建 micro 服务
	service := web.NewService(
		web.Name("go.micro.api.cinema"),
	)
	_ = service.Init()
	// 创建 gin 路由
	r := router.Router()
	s := http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 1 << 20,
	}

	// 用 gin 注册go-micro handler
	service.Handle("/", r)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("run gin error: %s", err)
	}
	if err := service.Run(); err != nil {
		log.Fatalf("run go-micro error: %s", err)
	}
}
