package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/model"
)

// handler layer 내 service 정의
type Handler struct {
	DeviceService model.DeviceService
	UserService   model.UserService
}

// 의존성이 주입되며 handler 레이어 초기설정
type Config struct {
	R             *gin.Engine
	UserService   model.UserService
	DeviceService model.DeviceService
}

func NewHandler(c *Config) {
	// 의존성 주입
	h := &Handler{
		UserService:   c.UserService,
		DeviceService: c.DeviceService,
	}

	v1 := c.R.Group("/v1")
	{
		deviceGroup := v1.Group("/device")
		{
			deviceGroup.POST("/register", h.DeviceRegister)
		}
		userGroup := v1.Group("/user")
		{
			userGroup.GET("/", h.UserGet)
			userGroup.POST("/login", h.Login)
		}
		subjectGroup := v1.Group("/subject")
		{
			subjectGroup.GET("/", h.Subject)
		}
	}
}

func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) Subject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}
