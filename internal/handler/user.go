package handler

import (
	. "bluebell/internal/params"
	"bluebell/internal/service"
	"bluebell/pkg/errno"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	service.Register()

	SendResponse(c, nil, nil)
}
