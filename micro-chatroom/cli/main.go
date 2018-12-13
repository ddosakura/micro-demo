package main

import (
	"context"
	"fmt"
	"log"

	upb "github.com/ddosakura/go-micro-demo/micro-chatroom/user-service/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	sercice := micro.NewService(micro.Name("go.micro.srv.user"))
	sercice.Init()

	client := upb.NewUserServiceClient("go.micro.srv.user", sercice.Client())

	fmt.Println("----- start -----")
	actionOne(&client)
	actionTwo(&client)
	actionOne(&client)
	fmt.Println("-----  end  -----")
}

func actionOne(client *upb.UserServiceClient) {
	rep, err := (*client).GetAll(context.Background(), &upb.Request{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rep)
}

func actionTwo(client *upb.UserServiceClient) {
	rep, err := (*client).Register(context.Background(), &upb.User{
		Username: "root",
		Password: "123456",
		Nickname: "ddosakura",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rep)
}
