package conf

import (
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

// [conf.Run()]
// Esta función recive la instancia del servidor,
// Primero valida si el entorno tiene configurado el servicio para https,
// Dependiendo de la variable, se inicia el servicio en http o https.
func Run(server *echo.Echo) {
	if sslCert, _ := strconv.ParseBool(os.Getenv("SSLCERT")); sslCert {
		server.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
		go httpRedirect()
		if err := server.StartAutoTLS(":443"); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := server.Start(":80"); err != nil {
			log.Fatal(err)
		}
	}
}

func httpRedirect() {
	httpRedirect := echo.New()
	httpRedirect.Pre(middleware.HTTPSRedirect())
	if err := httpRedirect.Start(":80"); err != nil {
		log.Fatal(err)
	}
}
