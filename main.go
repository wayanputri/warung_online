package main

import (
	"warung_online/app/config"
	"warung_online/app/database"
	"warung_online/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	mysql:=database.InitMysql(cfg)
	database.InittialMigration(mysql)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	router.InitRouter(mysql,e)
	e.Logger.Fatal(e.Start(":8080"))

}