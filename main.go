package main

import (
	"Api/root/services"

	"github.com/gin-gonic/gin"
)

// En la funcion main se encuentran todas las rutas que tiene nuestra api
// Que tipo de autentificacion ocuparan y su handler.
func main() {
	router := gin.Default()

	router.GET("/", services.Healthcheck)
	router.POST("/locationIdUnit", services.BasicAuth, services.LocationIdUnit)
	router.POST("/unitsAvailable", services.BasicAuth, services.UnitsAvailable)
	router.POST("/boroughs", services.BasicAuth, services.BoroughsAvailable)
	router.POST("/unitsBorough", services.BasicAuth, services.UnitsPerBorough)

	router.Run("0.0.0.0:3001")
}
