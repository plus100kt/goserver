package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/model/apperrors"
)

type loginReq struct {
	UUID     string `json:"uuid" binding: "required,uuid"`
	Key      string `json:"key" binding: "required,key"`
	ID       string `json:"id" binding: "required,id"`
	Password string `json:"password" binding: "required,password"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginReq
	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		UUID:        req.UUID,
		ID:          req.ID,
		AesPassword: req.Password,
	}

	u, err := h.UserService.Login(c, req.Key, u)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.Bind(u))
}
