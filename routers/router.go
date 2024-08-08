package routers

import (
	v1 "github.com/Vghxv/GinHub/routers/api/v1"

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
	router.GET(basePath+"/albums", v1.GetAlbums)
	router.GET(basePath+"/albums/:id", v1.GetAlbumByID)
	router.POST(basePath+"/albums", v1.PostAlbums)
	return router
}
