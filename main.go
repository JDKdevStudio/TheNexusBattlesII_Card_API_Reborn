package main

import (
	"TheNexusBattlesII_Card_API_Reborn/conf"

	_ "TheNexusBattlesII_Card_API_Reborn/docs"

	"github.com/labstack/echo/v4"
)

// @title           The Nexus Battles II: Card API Service Reborn
// @version         2.0
// @description     Servicio API para acceder a las cartas del negocio
// @termsOfService  http://swagger.io/terms/

// @license.name  MIT
// @license.url   https://opensource.org/license/mit/

// @host      localhost
// @BasePath  /api

// @externalDocs.description  Repository
// @externalDocs.url          https://github.com/JDKdevStudio/TheNexusBattlesII_Card_API_Reborn
func main() {
	server := echo.New()
	conf.Init(server)
	conf.Run(server)
}
