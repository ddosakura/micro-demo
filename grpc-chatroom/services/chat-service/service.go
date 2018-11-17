package chatService

import (
	"context"
	"log"
	"time"

	cpb "github.com/ddosakura/go-micro-demo/grpc-chatroom/proto/chat"
	"github.com/ddosakura/go-micro-demo/grpc-chatroom/utils"
)

type Service struct {
	Data []*cpb.ChatService_ListenServer
}

func (s *Service) Send(ctx context.Context, req *cpb.Request) (*cpb.Response, error) {
	claims, err := utils.DecodeJWT(req.GetToken())
	if err != nil {
		return &cpb.Response{}, err
	}

	for _, v := range s.Data {
		defer func() {
			_ = recover()
		}()
		(*v).Send(&cpb.Message{
			Nickname: claims.User.GetNickname(),
			Data:     req.GetMsg(),
		})
	}
	return &cpb.Response{}, nil
}

func (s *Service) Listen(req *cpb.Request, stream cpb.ChatService_ListenServer) error {
	claims, err := utils.DecodeJWT(req.GetToken())
	if err != nil {
		return err
	}

	nickname := claims.User.GetNickname()
	log.Println(nickname, "Login!")
	s.Data = append(s.Data, &stream)

	s.Send(context.Background(), &cpb.Request{
		Token: req.GetToken(),
		Msg:   "Login!",
	})

	for {
		time.Sleep(time.Second * 61)
		claims, err := utils.DecodeJWT(req.GetToken())
		if err != nil {
			return err
		}
		log.Println(claims)
	}
}
