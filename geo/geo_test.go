package geo_test

import (
	"testing"
	"weathercli/geo"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonsdfs"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrorNoCity, err)
	}
}
