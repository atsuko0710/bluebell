package handler

import (
	. "bluebell/internal/params"
	"bluebell/internal/service"
	"bluebell/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegisterHandler(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		// 判断错误是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			SendResponse(c, errno.ErrBind, nil)
			return
		}

		SendResponse(c, errno.ErrBind, removeTopStruct(errs.Translate(trans)))
		return
	}
	SendResponse(c, service.Register(r), nil)
}

func LoginHandler(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	
}
