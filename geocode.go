package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"bytes"
	"io/ioutil"
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
    url := "http://www.mapquestapi.com/geocoding/v1/reverse?key="+apiKey

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
		nerr := fmt.Errorf("%s: %s, No se pudo parsear, %s", err.Error(), body)
		fmt.Println((nerr.Error()))
		return
	}
	
	add.Boroungh=address.Results[0].Locations[0].AdminArea5
	add.City=address.Results[0].Locations[0].AdminArea4
	add.Street=address.Results[0].Locations[0].Street
	return add
}
