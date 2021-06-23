package repository

import (
	"bluebell/internal/models"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

// 检查用户名是否存在
func CheckUserExist(username string) bool {
	_, err := models.GetUserByName(username)
	// 没有在找到数据
	if err == gorm.ErrRecordNotFound {
		return true
	}
	if err != nil {
		log.Error("CheckUserExist", err)
		return false
	}
	return false
}

func CreateUser(u models.UserModel) error {
	return models.Create(u)
}

func GetUser(username string) (bool, models.UserModel) {
	user, err := models.GetUserByName(username)
	if err != nil {
		return false, models.UserModel{}
	}
	return true, *user
}