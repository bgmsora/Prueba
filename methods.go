package main

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

func reverseGeocode(lat float64, lng float64) (add responseAdressInterface) {
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

	add.Boroungh = address.Results[0].Locations[0].AdminArea5
	add.City = address.Results[0].Locations[0].AdminArea4
	add.Street = address.Results[0].Locations[0].Street
	return add
}

const queryBusId string = "{\"query\":\"query MyQuery ($id:Int){\\r\\n  mb(where: {vehicle_id: {_eq: $id}}) {\\r\\n    position_latitude\\r\\n    position_longitude\\r\\n  }\\r\\n}\",\"variables\":{\"id\":XWFFF}}"

// XWFFF

func hasuraRequestId(id int) (add responseHasuraId) {
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
