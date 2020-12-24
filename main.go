package main

import (
	"LtthStore/database"
	"LtthStore/middlewares"
	_ "LtthStore/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type", "Content-Type", "Auth"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, middlewares.Jwt)

	database.InitDatabase()
	beego.Run()
}
