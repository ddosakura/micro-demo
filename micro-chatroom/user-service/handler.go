package main

import (
	"context"
	"errors"

	upb "github.com/ddosakura/go-micro-demo/micro-chatroom/user-service/proto/user"
	"github.com/ddosakura/go-micro-demo/micro-chatroom/utils"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
	userRepo UserRepo
}

func (s *Service) Register(ctx context.Context,
	user *upb.User,
	res *upb.Response) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		res = &upb.Response{
			Code: -1,
		}
		return err
	}
	user.UserId = uuid.String()
	if err := s.userRepo.Insert(user); err != nil {
		res = &upb.Response{
			Code: -1,
		}
		return err
	}
	res = &upb.Response{
		Code: 0,
	}
	return nil
}

func (s *Service) Login(ctx context.Context,
	user *upb.User,
	res *upb.Response) error {
	u, err := s.userRepo.SelectByUsername(user.GetUsername())
	if err != nil {
		res = &upb.Response{
			Code: -1,
		}
		return err
	}
	if u.GetPassword() != user.GetPassword() {
		res = &upb.Response{
			Code: -1,
		}
		return errors.New("wrong password")
	}
	token, err := utils.EncodeJWT(u)
	if err != nil {
		res = &upb.Response{
			Code: -1,
		}
		return err
	}
	res = &upb.Response{
		Code:  0,
		Token: token,
	}
	return nil
}

func (s *Service) GetAll(ctx context.Context,
	req *upb.Request,
	res *upb.Response) error {
	l, err := s.userRepo.SelectAll()
	if err != nil {
		res = &upb.Response{
			Code: -1,
		}
		return err
	}
	res = &upb.Response{
		Code:  0,
		Users: l,
	}
	return nil

}
