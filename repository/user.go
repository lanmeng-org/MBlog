package repository

import (
	"MBlog/model"
	"MBlog/util"
	"errors"
)

func Login(name string, password string) (*model.User, error) {
	user, err := model.GetUserByName(name)
	if err != nil {
		return nil, err
	}

	if util.PwdIsRight(password, user.Password) {
		return user, nil
	}

	return nil, errors.New("User or password is error")
}