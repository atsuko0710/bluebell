package handler

import (
	"bluebell/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	service.Register()

	SendResponse(c, nil, nil)
}
