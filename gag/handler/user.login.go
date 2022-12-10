package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/model/app"
)

type loginReq struct {
	UUID     string `json:"uuid" binding: "required,uuid"`
	Key      string `json:"key" binding: "required,key"`
	ID       string `json:"id" binding: "required,id"`
	Password string `json:"password" binding: "required,password"`
}

type loginRes struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		UUID:        req.UUID,
		ID:          req.ID,
		AesPassword: req.Password,
	}

	err := h.UserService.Login(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(loginRes{
		Name:     user.Name,
		Email:    user.Email,
		ImageURL: user.ImageURL,
	})

	c.IndentedJSON(http.StatusOK, res)
}
