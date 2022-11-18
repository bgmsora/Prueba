package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Verifica los datos de conexion de un usuario
func BasicAuth(c *gin.Context) {
	reqUser, reqPass, hasAuth := c.Request.BasicAuth()

	var envUser string = os.Getenv("BASIC_AUTH_USER")
	var envPass string = os.Getenv("BASIC_AUTH_PASS")

	if hasAuth && (reqUser != envUser || reqPass != envPass) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}

// Obtiene la direccion de una unidad por su identificador unico (vehicle_id)
func LocationIdUnit(c *gin.Context) {
	// Body Http obtener
	var payload idInterface
	if err := c.BindJSON(&payload); err != nil {
		var errorString string = fmt.Sprintf("Error in payload: %v\n%s", payload, err.Error())
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}
	fmt.Println("Id consulted-> ", payload.Id)

	//Consulta a Graphql
	hasuraResponse := HasuraRequestId(payload.Id)
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Mb) == 0 {
		var errorString string = fmt.Sprintf("Id not exist in DB: %d\n", payload.Id)
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}

	//Obtener Direccion con base a las coordenadas y regresarla en formato Json
	response := ReverseGeocode(hasuraResponse.Data.Mb[0].PositionLatitude, hasuraResponse.Data.Mb[0].PositionLongitude)
	fmt.Println(response)
	bytesResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", bytesResponse)
}

// Obtienes las unidades disponibles
func UnitsAvailable(c *gin.Context) {
	//Consulta a Graphql
	hasuraResponse := HasuraRequestUnitAvailable()
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Mb) == 0 {
		var errorString string = fmt.Sprintf("No hay unidades disponibles")
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}
	bytesResponse, err := json.Marshal(hasuraResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", bytesResponse)
}

// Obtiene las alcaldias disponibles
func BoroughsAvailable(c *gin.Context) {
	//Consulta a Graphql
	hasuraResponse := HasuraRequestUnits()
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Mb) == 0 {
		var errorString string = fmt.Sprintf("No hay unidades disponibles")
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}

	//Obtener las alcaldias
	unitsBorough := getBorough(hasuraResponse)
	bytesResponse, err := json.Marshal(unitsBorough)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", bytesResponse)
}

// Obtiene las unidades con respecto a una alcaldia
func UnitsPerBorough(c *gin.Context) {
	// Body Http
	var payload boroughInterface
	if err := c.BindJSON(&payload); err != nil {
		var errorString string = fmt.Sprintf("Error in payload: %v\n%s", payload, err.Error())
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}
	fmt.Println("Borough consulted-> ", payload.Borough)

	//Consulta a Graphql
	hasuraResponse := HasuraRequestUnits()
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Mb) == 0 {
		var errorString string = fmt.Sprintf("No hay unidades disponibles")
		fmt.Println(errorString)
		c.Data(http.StatusOK, "application/json", []byte(errorString))
		return
	}

	//Filtro sobre la alcaldia
	unitsBorough := filterBorough(hasuraResponse, payload.Borough)
	bytesResponse, err := json.Marshal(unitsBorough)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", bytesResponse)
}

func Healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "Ok 200")
}
