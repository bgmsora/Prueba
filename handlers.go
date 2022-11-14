package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func basicAuth(c *gin.Context) {
	reqUser, reqPass, hasAuth := c.Request.BasicAuth()

	var envUser string = os.Getenv("BASIC_AUTH_USER")
	var envPass string = os.Getenv("BASIC_AUTH_PASS")

	if hasAuth && (reqUser != envUser || reqPass != envPass) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}

func LocationIdUnit(c *gin.Context) {
	//! Body Http
	var payload idInterface
	if err := c.BindJSON(&payload); err != nil {
		var errorString string = fmt.Sprintf("Error in payload: %v\n%s", payload, err.Error())
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}
	fmt.Println("Id consulted-> ", payload.Id)

	//Consulta a Hasura
	hasuraResponse := hasuraRequestId(payload.Id)
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Mb) == 0 {
		var errorString string = fmt.Sprintf("Id not exist in DB: %d\n", payload.Id)
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}

	//Obtener Direccion con base a las coordenadas y regresarla en formato Json
	response := reverseGeocode(hasuraResponse.Data.Mb[0].PositionLatitude, float64(hasuraResponse.Data.Mb[0].PositionLongitude))
	fmt.Println(response)
	bytesResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", bytesResponse)
}

func UnitsAvailable(c *gin.Context) {

}

func BoroughsAvailable(c *gin.Context) {

}

func unitsPerBorough(c *gin.Context) {

}

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "Ok 200")
}
