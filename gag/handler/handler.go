package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/model"
)

// handler layer 내 service 정의
type Handler struct {
	UserService model.UserService
}

// 의존성이 주입되며 handler 레이어 초기설정
type Config struct {
	R           *gin.Engine
	UserService model.UserService
}

func NewHandler(c *Config) {
	// 의존성 주입
	h := &Handler{
		UserService: c.UserService,
	}

	v1 := c.R.Group("/v1")
	{
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/login", h.Login)
			userGroup.POST("/register", h.DeviceRegister)
		}
	}
}
