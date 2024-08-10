package routers

import (
	album "github.com/Vghxv/GinHub/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// @title GinHub API
// @version 1.0
// @description GinHub API server
// @host localhost:8000
// @BasePath /api/v1
func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	basePath := "/api/v1"
	
	apiRouter := router.Group(basePath)
	{
		apiRouter.GET("/albums", album.GetAlbums)
		apiRouter.GET("/albums/:id", album.GetAlbumByID)
		apiRouter.POST("/albums", album.PostAlbums)
	}

	return router
}
