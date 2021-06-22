package service

import (
	"bluebell/internal/models"
	"bluebell/internal/params"
	"bluebell/internal/repository"

	"bluebell/pkg/errno"
	"bluebell/pkg/snowflake"

	"github.com/lexkong/log"
)

func Register(u params.CreateRequest) errno.Errno {

	if !repository.CheckUserExist(u.Username) {
		return *errno.ErrUserNameExist
	}

	userId := snowflake.GetId()

	user := models.UserModel{
		UserId: userId,
		Password: u.Password,
		Email: u.Email,
		UserName: u.Username,
		Gender: u.Gender,
	}
	if err := repository.CreateUser(user); err != nil {
		log.Error("RegisterService", err)
		return *errno.ErrCreateUser
	}
	return *errno.OK
}
