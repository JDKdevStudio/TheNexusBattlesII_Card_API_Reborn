package routes

import (
	"TheNexusBattlesII_Card_API_Reborn/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// [routes.Init()]
// Esta función define los endpoints para la aplicación,
// También injecta los controladores a cada ruta específica.
func Init(server *echo.Echo) {
	//Endpoints de utilidad
	server.GET("/docs/*", echoSwagger.WrapHandler)
	server.Static("/assets/", "assets")
	//Endpoints para heroe
	server.GET("/api/heroes/:id", controllers.GetHeroe)
	server.GET("/api/heroes/", controllers.GetHeroeList)
	server.GET("/api/heroes/name/", controllers.GetHeroeListByName)
	server.POST("/api/heroes/", controllers.PostHeroe)
	server.PATCH("api/heroes/", controllers.PatchHeroe)
}
