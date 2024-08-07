package main

import (
	"log"

	_ "github.com/Vghxv/GinHub/docs"

	"github.com/Vghxv/GinHub/database"
	"github.com/Vghxv/GinHub/pkg/setting"
	"github.com/Vghxv/GinHub/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GinHub API
// @version 1
// @description This is a GinHub API server.
// @BasePath /api/v1
func main() {
	setting.Init()
	database.Connect()
	r := routers.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	host := setting.ServerSetting.Host
	port := setting.ServerSetting.Port
	log.Fatal(r.Run(host + ":" + port))
}
