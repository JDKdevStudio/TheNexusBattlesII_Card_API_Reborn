package main

import (
	"TheNexusBattlesII_Card_API_Reborn/conf"

	"github.com/labstack/echo/v4"
)

// [main()]
// Función de arranque de la aplicación,
// Primero instancia un servidor nuevo,
// Luego inicia la configuración preliminar,
// Por último inicia el servicio según los parámetros dados.
func main() {
	server := echo.New()
	conf.Init(server)
	conf.Run(server)
}
