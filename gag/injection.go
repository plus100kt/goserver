package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/plus100kt/goserver/gag/eclass"
	"github.com/plus100kt/goserver/gag/handler"
	"github.com/plus100kt/goserver/gag/repository"
	"github.com/plus100kt/goserver/gag/service"
)

func inject(d *dataSources) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	/*
	 * repository layer
	 */
	userRepository := repository.NewUserRepository(d.DB)
	deviceRepository := repository.NewDeviceRepository(d.DB)
	eclassRepository := repository.NewEclassRepository(&repository.EclassConfg{})

	/*
	 * service layer
	 */
	userService := service.NewUserService(&service.USConfig{
		UserRepository:   userRepository,
		DeviceRepository: deviceRepository,
		EclassRepository: eclassRepository,
	})

	router := gin.Default()

	eclass.NewEclass(&eclass.EclasssConfig{})
	handler.NewHandler(&handler.Config{
		R:           router,
		UserService: userService,
	})

	return router, nil
}
