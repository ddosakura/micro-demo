package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/ddosakura/go-micro-demo/micro-chatroom/foo/handler"
	"github.com/ddosakura/go-micro-demo/micro-chatroom/foo/subscriber"

	example "github.com/ddosakura/go-micro-demo/micro-chatroom/foo/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.foo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.foo", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.foo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
