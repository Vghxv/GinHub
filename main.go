package main

import (
	"log"

	"github.com/Vghxv/GinHub/database"
	"github.com/Vghxv/GinHub/pkg/setting"
	"github.com/Vghxv/GinHub/routers"
)

func main() {
	setting.Init()
	database.Connect()
	r := routers.SetupRouter()
	host := setting.ServerSetting.Host
	port := setting.ServerSetting.Port
	log.Fatal(r.Run(host + ":" + port))
}
