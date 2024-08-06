package router

import (
	"github.com/Vghxv/GinHub/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)

	return router
}
