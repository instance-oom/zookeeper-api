package main

import (
	"zookeeper-api/config"
	"zookeeper-api/controllers"

	"github.com/astaxie/beego/plugins/cors"

	"zookeeper-api/routers"

	"github.com/astaxie/beego"
)

func main() {

	config.LoadConfig()
	controllers.Init()
	routers.Init()

	beego.BConfig.Listen.HTTPPort = config.ListenPort
	beego.SetLevel(config.LogLevel)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	beego.Run()
}
