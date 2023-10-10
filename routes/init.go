package routes

import (
	"TheNexusBattlesII_Card_API_Reborn/controllers"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// [routes.Init()]
// Esta función define los endpoints para la aplicación,
// También injecta los controladores a cada ruta específica.
func Init(server *echo.Echo) {
	//Endpoints de utilidad
	server.GET("/", controllers.RedirectToDocs)
	server.GET("/docs/*", echoSwagger.WrapHandler)
	server.Static("/assets/", os.Getenv("ASSETSFOLDER"))
	//Endpoints de uso general
	server.GET("/api/cartas/:id", controllers.GetObject)
	server.GET("/api/cartas/", controllers.GetObjectList)
	//Endpoints para heroe
	server.POST("/api/heroes/", controllers.PostHeroe)
	server.PATCH("api/heroes/", controllers.PatchHeroe)
	//Endpoints para consumible
	server.POST("/api/consumible/", controllers.PostConsumible)
	server.PATCH("/api/consumible/", controllers.PatchConsumible)
}
