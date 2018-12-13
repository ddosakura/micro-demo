package main

import (
	"errors"

	upb "github.com/ddosakura/go-micro-demo/micro-chatroom/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	Insert(*upb.User) error
	SelectByUsername(string) (*upb.User, error)
	SelectAll() ([]string, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func (repo *UserRepoImpl) Insert(u *upb.User) error {
	if _, err := repo.SelectByUsername(u.GetUsername()); err == nil {
		return errors.New("user has exist!")
	}
	if err := repo.db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepoImpl) SelectByUsername(username string) (*upb.User, error) {
	u := &upb.User{
		Username: username,
	}
	if err := repo.db.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepoImpl) SelectAll() ([]string, error) {
	var users []*upb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	ret := make([]string, len(users))
	for i := range users {
		ret[i] = users[i].GetNickname()
	}
	return ret, nil
}
