package router

import (
	"songs/config"
	"songs/controller"
	"songs/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = ""

	songs := router.Group("/songs")
	{
		songs.POST("", controller.CreateSong)
		songs.DELETE("", controller.DeleteSong)
		songs.POST("/find", controller.FindSongs)
		songs.PUT("", controller.UpdateSong)

		songs.POST("/text", controller.GetText)
	}

	if config.MockApi {
		router.GET("/info", controller.GetInfo)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
