package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/model/apperrors"
)

type deviceRegisterReq struct {
	UUID string `json:"uuid" binding: "required,uuid"`
}

func (h *Handler) DeviceRegister(c *gin.Context) {
	var req deviceRegisterReq
	if ok := bindData(c, &req); !ok {
		return
	}

	uuid := req.UUID
	public_key, err := h.UserService.DeviceRegister(c, uuid)
	if err != nil {
		log.Printf("register error: %v\n%v", uuid, err)
		e := apperrors.NewInternal()

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"public_key": public_key,
	})
}
