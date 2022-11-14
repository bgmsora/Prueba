package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", healthcheck)
	router.POST("/locationIdUnit", basicAuth, LocationIdUnit)
	router.POST("/locationsUnits", basicAuth, LocationsUnits)
	router.POST("/boroughs", basicAuth, BoroughsAvailable)
	router.POST("/unitsBorough", basicAuth, unitsPerBorough)

	router.Run("0.0.0.0:3001")
}