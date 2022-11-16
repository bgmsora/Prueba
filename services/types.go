package services

type idInterface struct {
	Id int `json:"id"`
}

type responseAdressInterface struct {
	Boroungh string `json:"alcaldia"`
	City     string `json:"ciudad"`
	Street   string `json:"calle"`
}

type Info struct {
	Copyright struct {
		Text         string `json:"text"`
		ImageURL     string `json:"imageUrl"`
		ImageAltText string `json:"imageAltText"`
	} `json:"copyright"`
	Statuscode int      `json:"statuscode"`
	Messages   []string `json:"messages"`
}

type GeocodingResult struct {
	Info    Info `json:"info"`
	Options struct {
		MaxResults        int  `json:"maxResults"`
		ThumbMaps         bool `json:"thumbMaps"`
		IgnoreLatLngInput bool `json:"ignoreLatLngInput"`
	} `json:"options"`
	Results []struct {
		ProvidedLocation struct {
			Location string `json:"location"`
		} `json:"providedLocation"`
		Locations []struct {
			Street string `json:"street"`
			// Neighborhood
			AdminArea6     string `json:"adminArea6"`
			AdminArea6Type string `json:"adminArea6Type"`
			// City
			AdminArea5     string `json:"adminArea5"`
			AdminArea5Type string `json:"adminArea5Type"`
			// County
			AdminArea4     string `json:"adminArea4"`
			AdminArea4Type string `json:"adminArea4Type"`
			// State
			AdminArea3     string `json:"adminArea3"`
			AdminArea3Type string `json:"adminArea3Type"`
			// Country
			AdminArea1         string `json:"adminArea1"`
			AdminArea1Type     string `json:"adminArea1Type"`
			PostalCode         string `json:"postalCode"`
			GeocodeQualityCode string `json:"geocodeQualityCode"`
			// ex: "NEIGHBORHOOD", "CITY", "COUNTY"
			GeocodeQuality string `json:"geocodeQuality"`
			DragPoint      bool   `json:"dragPoint"`
			SideOfStreet   string `json:"sideOfStreet"`
			LinkId         string `json:"linkId"`
			UnknownInput   string `json:"unknownInput"`
			Type           string `json:"type"`
			LatLng         LatLng `json:"latLng"`
			DisplayLatLng  LatLng `json:"displayLatLng"`
			MapUrl         string `json:"mapUrl"`
		} `json:"locations"`
	} `json:"results"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type locationStructure struct {
	Location struct {
		LatLng struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"latLng"`
	} `json:"location"`
	Options struct {
		ThumbMaps bool `json:"thumbMaps"`
	} `json:"options"`
	IncludeNearestIntersection bool `json:"includeNearestIntersection"`
	IncludeRoadMetadata        bool `json:"includeRoadMetadata"`
}

type responseHasuraId struct {
	Data struct {
		Mb []struct {
			PositionLatitude  float64 `json:"position_latitude"`
			PositionLongitude int     `json:"position_longitude"`
		} `json:"mb"`
	} `json:"data"`
}

type responseHasuraUnitsAvailable struct {
	Data struct {
		Mb []struct {
			dataVehicle
		} `json:"mb"`
	} `json:"data"`
}

type dataVehicle struct {
	Vehicle_id        int     `json:"vehicle_id"`
	PositionLatitude  float64 `json:"position_latitude"`
	PositionLongitude int     `json:"position_longitude"`
	Trip_start_date   int     `json:"trip_start_date"`
	Trip_id           int     `json:"trip_id"`
	Position_speed    int     `json:"position_speed"`
}
