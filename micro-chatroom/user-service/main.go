package main

import (
	"log"

	upb "github.com/ddosakura/go-micro-demo/micro-chatroom/user-service/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	defer db.Close()

	// 自动检查 User 结构是否变化
	db.AutoMigrate(&upb.User{})
	repo := &UserRepoImpl{db}

	server := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	server.Init()

	upb.RegisterUserServiceHandler(server.Server(), &Service{
		userRepo: repo,
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
