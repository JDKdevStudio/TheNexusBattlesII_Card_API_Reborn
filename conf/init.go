package conf

import (
	"TheNexusBattlesII_Card_API_Reborn/env"
	"TheNexusBattlesII_Card_API_Reborn/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// [conf.Init()]
// Esta función recive la instancia del servidor,
// Primero configura las variables de entorno,
// Luego configura el CORS mediante un middleware,
// Después manda a llamar la configuración de rutas para los endpoints.
func Init(server *echo.Echo) {
	env.Init()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch},
	}))
	routes.Init(server)
}
