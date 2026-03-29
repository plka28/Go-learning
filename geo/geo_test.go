package geo_test

import (
	"demo/app-5/geo"
	"testing"
)

func TestGetNyLocation(t *testing.T) {
	//Arrange
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	//Act
	got, err := geo.GetMyLocation(city)
	//Assert
	if err != nil {
		t.Error("Ошибка получения города")
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}
}
