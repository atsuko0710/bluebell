package models

import (
	"bluebell/config"
)

type UserModel struct {
	BaseModel
	UserId   int64  `json:"user_id" gorm:"column:user_id;not null"`
	UserName string `json:"username" gorm:"column:username;not null" bindding:"required"`
	Password string `json:"password" gorm:"column:password;not null" bindding:"required"`
	Email    string `json:"email" gorm:"column:email;not null" bindding:"required"`
	Gender   int    `json:"gender" gorm:"column:gender;not null" bindding:"required"`
}

func (u *UserModel) TableName() string {
	return "user"
}

func GetUserByName(username string) (*UserModel, error) {
	u := &UserModel{}
	d := config.DB.Self.Where("username=?", username).First(&u)
	return u, d.Error
}

func Create(u UserModel) error {
	return config.DB.Self.Create(&u).Error
}
