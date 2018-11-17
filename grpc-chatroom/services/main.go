package main

import (
	"log"
	"net"

	cpb "github.com/ddosakura/go-micro-demo/grpc-chatroom/proto/chat"
	upb "github.com/ddosakura/go-micro-demo/grpc-chatroom/proto/user"
	"github.com/ddosakura/go-micro-demo/grpc-chatroom/services/chat-service"
	"github.com/ddosakura/go-micro-demo/grpc-chatroom/services/user-service"
	"google.golang.org/grpc"
)

const (
	PORT = ":50001"
)

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on: %s\n", PORT)

	server := grpc.NewServer()

	// 向 rRPC 服务器注册微服务
	upb.RegisterUserServiceServer(
		server,
		&userService.Service{
			UserDao: &userService.UserDaoImpl{},
		},
	)
	cpb.RegisterChatServiceServer(
		server,
		&chatService.Service{
			Data: make([]*cpb.ChatService_ListenServer, 0),
		},
	)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
