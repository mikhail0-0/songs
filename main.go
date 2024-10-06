package main

import (
	"fmt"
	"songs/config"
	"songs/postgres"
	"songs/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Songs Library API
//	@version		1.0
//	@description	API for library of songs
func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	_, err = postgres.GetDB()
	if err != nil {
		panic(err)
	}

	router := router.GetRouter()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(fmt.Sprintf(":%v", config.AppPort))
}
