package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	//"github.com/hasura/go-graphql-client" //para graphql
	//"golang.org/x/oauth2" //para graphql
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
		return
	}
	fmt.Println("Id consulted-> ", payload.Id)

	//Consulta a Hasura
	/*
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient("http://localhost:8080/graphql", httpClient)
    */

	//Obtener Direccion con base a las coordenadas y regresarla en formato Json
	response := reverseGeocode(19.440222, -99.133207)
	fmt.Println(response)
	bytesResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", bytesResponse)
}


func LocationsUnits(c *gin.Context) {

}

func BoroughsAvailable(c *gin.Context) {

}

func unitsPerBorough(c *gin.Context) {
	
}

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "Ok 200")
}
