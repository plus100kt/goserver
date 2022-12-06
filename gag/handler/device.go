package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/model/apperrors"
)

func (h *Handler) DeviceRegister(c *gin.Context) {
	device, exists := c.Get("device")
	if !exists {
		log.Printf("request 로부터 device 추출 실패: %v\n", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}
	uuid := device.(*model.Device).UUID

	public_key, err := h.DeviceService.Register(c, uuid)
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
