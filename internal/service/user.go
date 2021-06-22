package service

import (
	"bluebell/internal/models"
	"bluebell/internal/params"
	"bluebell/internal/repository"

	"bluebell/pkg/auth"
	"bluebell/pkg/errno"
	"bluebell/pkg/snowflake"

	"github.com/lexkong/log"
)

func Register(u params.CreateRequest) errno.Errno {

	if !repository.CheckUserExist(u.Username) {
		return *errno.ErrUserNameExist
	}

	userId := snowflake.GetId()

	pass, err := auth.Encrypt(u.Password)
	if err != nil {
		return *errno.ErrCreateUser
	}

	user := models.UserModel{
		UserId:   userId,
		Password: pass,
		Email:    u.Email,
		UserName: u.Username,
		Gender:   u.Gender,
	}
	if err := repository.CreateUser(user); err != nil {
		log.Error("RegisterService", err)
		return *errno.ErrCreateUser
	}
	return *errno.OK
}

func Login(u params.LoginRequest) {
	
}
