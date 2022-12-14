package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/model/app"
)

type deviceRegisterReq struct {
	UUID string `json:"uuid" binding: "required,uuid"`
}

type deviceRegisterRes struct {
	PublicKey string `json:"public_key"`
}

func (h *Handler) DeviceRegister(c *gin.Context) {
	var req deviceRegisterReq
	if ok := bindData(c, &req); !ok {
		return
	}

	uuid := req.UUID
	device, err := h.UserService.DeviceRegister(c, uuid)
	if err != nil {
		log.Printf("register error: %v\n%v", uuid, err)
		e := app.NewInternal()

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	res := app.NewSuccess(deviceRegisterRes{
		PublicKey: device.RsaPublicKey,
	})

	c.IndentedJSON(http.StatusOK, res)
}
