package main

import (
	"bufio"
	"context"
	"log"
	"os"

	cpb "github.com/ddosakura/go-micro-demo/grpc-chatroom/proto/chat"
	upb "github.com/ddosakura/go-micro-demo/grpc-chatroom/proto/user"
	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:50001"
)

type ClientGroup struct {
	token string
	user  upb.UserServiceClient
	chat  cpb.ChatServiceClient
}

var (
	conn    *grpc.ClientConn
	clients ClientGroup
)

func init() {
	var err error
	conn, err = grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
}

func main() {
	defer conn.Close()
	defer func() {
		e := recover()
		if e != nil {
			log.Fatalln(e)
		}
	}()

	var cmd string
	var u upb.User
	cmd, u.Username, u.Password = os.Args[1], os.Args[2], os.Args[3]

	clients.user = upb.NewUserServiceClient(conn)

	switch cmd {
	case "r":
		u.Nickname = os.Args[4]
		resp, err := clients.user.Register(context.Background(), &u)
		if err != nil {
			panic(err.Error())
		}
		if resp.GetCode() != 0 {
			panic("r: unknow error!")
		}
		fallthrough
	case "l":
		resp, err := clients.user.Login(context.Background(), &u)
		if err != nil {
			panic(err.Error())
		}
		if resp.GetCode() != 0 {
			panic("l: unknow error!")
		}
		clients.token = resp.GetToken()
		gotoTalk()
	default:
		panic("please choose r/l")
	}
}

func gotoTalk() {
	resp, err := clients.user.GetAll(context.Background(), &upb.Request{})
	if err != nil {
		panic(err.Error())
	}
	if resp.GetCode() != 0 {
		panic("l: unknow error!")
	}

	log.Println(resp.GetUsers())

	//  chat client
	clients.chat = cpb.NewChatServiceClient(conn)

	stream, err := clients.chat.Listen(context.Background(), &cpb.Request{
		Token: clients.token,
	})
	if err != nil {
		panic(err)
	}

	go func() {
		// receive chat
		for {
			resp, err := stream.Recv()
			if err != nil {
				panic(err)
			}
			log.Println("[", resp.GetNickname(), "]:", resp.GetData())
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	for true {
		if ok := input.Scan(); ok {
			msg := input.Text()
			// fmt.Println(msg)

			// send chat
			_, err := clients.chat.Send(context.Background(), &cpb.Request{
				Token: clients.token,
				Msg:   msg,
			})
			if err != nil {
				panic(err)
			}
		} else {
			break
		}
	}
}
