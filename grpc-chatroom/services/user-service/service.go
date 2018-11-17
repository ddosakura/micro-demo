package userService

import (
	"context"
	"errors"

	upb "github.com/ddosakura/go-micro-demo/grpc-chatroom/proto/user"
	"github.com/ddosakura/go-micro-demo/grpc-chatroom/utils"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
	UserDao UserDao
}

func (s *Service) Register(ctx context.Context, user *upb.User) (*upb.Response, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return &upb.Response{
			Code: -1,
		}, err
	}
	user.UserId = uuid.String()
	if ok := s.UserDao.Insert(user); ok {
		return &upb.Response{
			Code: 0,
		}, nil
	}
	return &upb.Response{
		Code: -1,
	}, errors.New("username has exist!")
}

func (s *Service) Login(ctx context.Context, user *upb.User) (*upb.Response, error) {
	u := s.UserDao.SelectByUsername(user.GetUsername())
	if u == nil {
		return &upb.Response{
			Code: -1,
		}, errors.New("unknow user")
	}
	if u.GetPassword() != user.GetPassword() {
		return &upb.Response{
			Code: -1,
		}, errors.New("wrong password")
	}
	token, err := utils.EncodeJWT(u)
	if err != nil {
		return &upb.Response{
			Code: -1,
		}, err
	}
	return &upb.Response{
		Code:  0,
		Token: token,
	}, nil
}

func (s *Service) GetAll(ctx context.Context, req *upb.Request) (*upb.Response, error) {
	l := s.UserDao.SelectAll()
	return &upb.Response{
		Code:  0,
		Users: l,
	}, nil
}
