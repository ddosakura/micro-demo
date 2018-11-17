package userService

import (
	upb "github.com/ddosakura/go-micro-demo/grpc-chatroom/proto/user"
)

type UserDao interface {
	Insert(*upb.User) bool
	SelectByUsername(string) *upb.User
	SelectAll() []string
}

type UserDaoImpl struct {
	data []*upb.User
}

func (impl *UserDaoImpl) Insert(user *upb.User) bool {
	t := impl.SelectByUsername(user.GetUsername())
	if t == nil {
		impl.data = append(impl.data, user)
		return true
	}
	return false
}

func (impl *UserDaoImpl) SelectByUsername(username string) *upb.User {
	for i := range impl.data {
		if impl.data[i].GetUsername() == username {
			return impl.data[i]
		}
	}
	return nil
}

func (impl *UserDaoImpl) SelectAll() (ret []string) {
	ret = make([]string, len(impl.data))
	for i := range impl.data {
		ret[i] = impl.data[i].GetNickname()
	}
	return
}
