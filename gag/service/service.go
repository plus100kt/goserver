package service

import (
	"github.com/plus100kt/goserver/gag/model"
)

type userService struct {
	UserRepository   model.UserRepository
	DeviceRepository model.DeviceRepository
	EclassRepository model.EclassRepository
}

type USConfig struct {
	UserRepository   model.UserRepository
	DeviceRepository model.DeviceRepository
	EclassRepository model.EclassRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository:   c.UserRepository,
		DeviceRepository: c.DeviceRepository,
		EclassRepository: c.EclassRepository,
	}
}
