package service

import (
	"bluebell/internal/repository"
	"bluebell/pkg/snowflake"
)

func Register() {

	repository.QueryUserByName()

	snowflake.GetId()

	repository.CreateUser()
}
