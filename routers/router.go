package routers

import (
	v1 "github.com/Vghxv/GinHub/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/albums", v1.GetAlbums)
	router.GET("/albums/:id", v1.GetAlbumByID)
	router.POST("/albums", v1.PostAlbums)

	return router
}
