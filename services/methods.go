package services

import (
	"Api/root/services/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Funcion que nos regresa una direccion apartir de la latitud y longitud
func ReverseGeocode(lat float64, lng float64) (addressResponse responseAdressInterface) {
	var data locationStructure
	data.Location.LatLng.Lat = lat
	data.Location.LatLng.Lng = lng
	data.Options.ThumbMaps = false
	data.IncludeNearestIntersection = true
	data.IncludeRoadMetadata = true
	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	var apiKey string = os.Getenv("API_GEOCODING")
	url := "http://www.mapquestapi.com/geocoding/v1/reverse?key=" + apiKey

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var address GeocodingResult
	if err := json.Unmarshal(body, &address); err != nil {
		nerr := fmt.Errorf("%s: %s, No se pudo parsear", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}

	//Si  la direccion no es valida
	if len(address.Results[0].Locations) == 0 {
		return
	}

	addressResponse.Borough = address.Results[0].Locations[0].AdminArea5
	addressResponse.City = address.Results[0].Locations[0].AdminArea4
	addressResponse.Street = address.Results[0].Locations[0].Street
	return addressResponse
}

// Solicitud a graphql de la posicion de latitud y longitud con respecto al identificador
func HasuraRequestId(id int) (position responseHasuraId) {
	data := strings.Replace(tools.QueryBusId, "XWFFF", strconv.Itoa(id), 1)
	fmt.Println(data)
	payload := strings.NewReader(data)

	body := hasuraRequest(payload)
	if body == nil {
		fmt.Println(body)
		return
	}
	fmt.Println(string(body))

	if err := json.Unmarshal(body, &position); err != nil {
		nerr := fmt.Errorf("%s: %s, No se pudo parsear", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}
	return position
}

// Solitud de graphql para obtener los datos necesarios de todas las unidades disponibles
func HasuraRequestUnitAvailable() (vehicles responseHasuraUnitsAvailable) {
	payload := strings.NewReader(tools.QueryUnitsAvailable)
	body := hasuraRequest(payload)
	if body == nil {
		fmt.Println(body)
		return
	}
	fmt.Println(string(body))

	if err := json.Unmarshal(body, &vehicles); err != nil {
		nerr := fmt.Errorf("%s: %s, No se pudo parsear", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}
	return vehicles
}

// Solicitud de graphql para obtener todos los datos necesarios de todas las unidades
func HasuraRequestUnits() (vehicles responseHasuraUnitsAvailable) {
	payload := strings.NewReader(tools.QueryUnits)
	body := hasuraRequest(payload)
	if body == nil {
		fmt.Println(body)
		return
	}
	fmt.Println(string(body))

	if err := json.Unmarshal(body, &vehicles); err != nil {
		nerr := fmt.Errorf("%s: %s, No se pudo parsear", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}
	return vehicles
}

// Se hace un filtro sobre las direcciones para obtener las alcaldias
func getBorough(units responseHasuraUnitsAvailable) []string {
	boroughs := make([]string, 31)
	for _, arr := range units.Data.Mb {
		response := ReverseGeocode(arr.PositionLatitude, arr.PositionLongitude)
		fmt.Println(response)
		if response.City == "Ciudad de México" {
			boroughs = append(boroughs, response.Borough)
		}
	}
	aux := RemoveDuplicateStr(boroughs)
	remove := aux[1:]
	return remove
}

// Filtro para obtener solo las unidades de determinada alcaldia
func filterBorough(units responseHasuraUnitsAvailable, borough string) (vehicles responseHasuraUnitsAvailable) {
	for _, arr := range units.Data.Mb {
		response := ReverseGeocode(arr.PositionLatitude, arr.PositionLongitude)
		fmt.Println(response)
		if response.Borough == borough {
			vehicles.Data.Mb = append(vehicles.Data.Mb, arr)
		}
	}
	return vehicles
}

// Funcion generica para hacer la comunicacion con graphql
func hasuraRequest(payload *strings.Reader) (body []byte) {
	url := "http://host.docker.internal:8080/v1/graphql"
	method := "POST"
	var err error
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	var hasuraSecret string = os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	req.Header.Add("x-hasura-admin-secret", hasuraSecret)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return body
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
