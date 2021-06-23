package service

import (
	"bluebell/internal/models"
	"bluebell/internal/params"
	"bluebell/internal/repository"

	"bluebell/pkg/auth"
	"bluebell/pkg/errno"
	"bluebell/pkg/snowflake"
	"bluebell/pkg/token"

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

func Login(u params.LoginRequest) (errno.Errno, string) {
	res, user := repository.GetUser(u.Username)
	if !res {
		return *errno.ErrUserNotFound, ""
	}

	// 比较数据库密码和传输来的密码
	if err := auth.Compare(user.Password, u.Password); err != nil {
		return *errno.ErrPasswordIncorrect, ""
	}

	t, err := token.Sign(token.Context{
		ID:       user.Id,
		Username: user.UserName,
	}, "")
	if err != nil {
		return *errno.ErrToken, ""
	}

	return *errno.OK, t
}
