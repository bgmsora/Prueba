package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Funcion que nos regesa una direccion apartir de la latitud y longitud
func ReverseGeocode(lat float64, lng float64) (add responseAdressInterface) {
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

	//Si retorna vacio, porque la direccion no es valida
	if len(address.Results[0].Locations) == 0 {
		return
	}
	add.Borough = address.Results[0].Locations[0].AdminArea5
	add.City = address.Results[0].Locations[0].AdminArea4
	add.Street = address.Results[0].Locations[0].Street
	return add
}

const queryBusId string = "{\"query\":\"query MyQuery ($id:Int){\\r\\n  mb(where: {vehicle_id: {_eq: $id}}) {\\r\\n    position_latitude\\r\\n    position_longitude\\r\\n  }\\r\\n}\",\"variables\":{\"id\":XWFFF}}"

func HasuraRequestId(id int) (add responseHasuraId) {
	data := strings.Replace(queryBusId, "XWFFF", strconv.Itoa(id), 1)
	fmt.Println(data)
	payload := strings.NewReader(data)

	body := hasuraRequest(payload)
	if body == nil {
		fmt.Println(body)
		return
	}
	fmt.Println(string(body))

	if err := json.Unmarshal(body, &add); err != nil {
		nerr := fmt.Errorf("%s: %s, No se pudo parsear", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}
	return add
}

const queryUnitsAvailable string = "{\"query\":\"query MyQuery {\\r\\n  mb(where: {trip_schedule_relationship: {_eq: 0}}) {\\r\\n    vehicle_id\\r\\n    position_latitude\\r\\n    position_longitude\\r\\n    trip_start_date\\r\\n    trip_id\\r\\n    position_speed\\r\\n    \\r\\n  }\\r\\n}\",\"variables\":{}}"

func HasuraRequestUnitAvailable() (add responseHasuraUnitsAvailable) {
	payload := strings.NewReader(queryUnitsAvailable)
	body := hasuraRequest(payload)
	if body == nil {
		fmt.Println(body)
		return
	}
	fmt.Println(string(body))

	if err := json.Unmarshal(body, &add); err != nil {
		nerr := fmt.Errorf("%s: %s, No se pudo parsear", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}
	return add
}

const queryUnits string = "{\"query\":\"query MyQuery {\\r\\n  mb {\\r\\n    vehicle_id\\r\\n    position_latitude\\r\\n    position_longitude\\r\\n    trip_start_date\\r\\n    trip_id\\r\\n    position_speed\\r\\n    \\r\\n  }\\r\\n}\",\"variables\":{}}"

func HasuraRequestUnits() (add responseHasuraUnitsAvailable) {
	payload := strings.NewReader(queryUnits)
	body := hasuraRequest(payload)
	if body == nil {
		fmt.Println(body)
		return
	}
	fmt.Println(string(body))

	if err := json.Unmarshal(body, &add); err != nil {
		nerr := fmt.Errorf("%s: %s, No se pudo parsear", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}
	return add
}

func filterBorough(units responseHasuraUnitsAvailable, borough string) (new responseHasuraUnitsAvailable) {
	for _, arr := range units.Data.Mb {
		response := ReverseGeocode(arr.PositionLatitude, arr.PositionLongitude)
		fmt.Println(response)
		if response.Borough == borough {
			new.Data.Mb = append(new.Data.Mb, arr)
		}
	}
	return new
}

func getBorough(units responseHasuraUnitsAvailable) []string {
	boroughs := make([]string, 31)
	for _, arr := range units.Data.Mb {
		response := ReverseGeocode(arr.PositionLatitude, arr.PositionLongitude)
		fmt.Println(response)
		if response.City == "Ciudad de México" {
			boroughs = append(boroughs, response.Borough)
		}
	}
	aux := removeDuplicateStr(boroughs)
	remove := aux[1:]
	return remove
}

func removeDuplicateStr(strSlice []string) []string {
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
