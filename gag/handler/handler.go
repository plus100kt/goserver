package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct{}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{} // currently has no properties
	v1 := c.R.Group("/v1")
	{
		v1.POST("/device/register", h.DeviceRegister)
		userGroup := v1.Group("/user")
		{
			userGroup.GET("/", h.User)
			userGroup.POST("/login", h.Login)
		}
		subjectGroup := v1.Group("/subject")
		{
			subjectGroup.GET("/", h.Subject)
		}
	}
}

func (h *Handler) User(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) DeviceRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signup",
	})
}

func (h *Handler) Subject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}
