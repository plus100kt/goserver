package main

import (
	"fmt"

	"github.com/plus100kt/goserver/gag/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dataSources struct {
	DB *gorm.DB
}

func initDS() (*dataSources, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Device{})

	return &dataSources{
		DB: db,
	}, nil
}
