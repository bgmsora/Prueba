package services_test

import (
	"Api/root/services"
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetAddress(t *testing.T) {
	//En las pruebas unitarias de Go, se debe anexar las variables de ambiente manualmente
	t.Setenv("API_GEOCODING", "pAGhvIdAnzlQOKMEwl6jK7zWygedjlYG")

	testCases := []struct {
		Name     string
		Lat      float64
		Lng      float64
		Expected error
	}{
		{
			Name:     "Coordenada sin datos",
			Lat:      99.12,
			Lng:      45.1,
			Expected: nil,
		},
		{
			Name:     "Coordenada CDMX Gustavo A. Madero",
			Lat:      19.519730830758117,
			Lng:      -99.15811643815897,
			Expected: nil,
		},
		{
			Name:     "Coordenada CDMX Venustiano Carranza",
			Lat:      19.413677117124394,
			Lng:      -99.12800091112778,
			Expected: nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel() //Crea hilos para cada solicitud

			result := services.ReverseGeocode(tc.Lat, tc.Lng)

			_, err := json.Marshal(result)
			if err == tc.Expected {
				t.Log("Test finalizado correctamente ", result, err)
			} else {
				t.Error("La respuesta tiene un json no valido")
			}
		})
	}
}

func TestHasuraRequestIdVehicle(t *testing.T) {
	//En las pruebas unitarias de Go, se debe anexar las variables de ambiente manualmente
	t.Setenv("HASURA_GRAPHQL_ADMIN_SECRET", "123456")

	testCases := []struct {
		Name     string
		Id       int
		Expected error
	}{
		{
			Name:     "Id no registrado",
			Id:       1,
			Expected: nil,
		},
		{
			Name:     "Id registrado 2",
			Id:       2,
			Expected: nil,
		},
		{
			Name:     "Id registrado 3",
			Id:       3,
			Expected: nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			result := services.HasuraRequestId(tc.Id)

			if len(result.Data.Mb) == 0 {
				t.Log("Id not exist in DB: ", tc.Id)
			} else {
				t.Log("Test finalizado correctamente ", result)
			}
		})
	}
}

func TestRemoveDuplicate(t *testing.T) {
	testCases := []struct {
		Name     string
		Slice    []string
		Expected []string
	}{
		{
			Name:     "Id no registrado",
			Slice:    []string{"India", "Canada", "Japan", "India", "Canada"},
			Expected: []string{"India", "Canada", "Japan"},
		},
		{
			Name:     "Id registrado 2",
			Slice:    []string{"a", "a", "a", "a", "a"},
			Expected: []string{"a"},
		},
		{
			Name:     "Id registrado 3",
			Slice:    []string{"Iztapalapa", "Iztapalapa", "Tláhuac", "Tláhuac", "Gustavo A. Madero"},
			Expected: []string{"Iztapalapa", "Tláhuac", "Gustavo A. Madero"},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			result := services.RemoveDuplicateStr(tc.Slice)

			if reflect.DeepEqual(result, tc.Expected) {
				t.Log("Test finalizado correctamente ", result)
			} else {
				t.Error("Test no pasado ", result, " != ", tc.Expected)
			}
		})
	}
}
