package types

import (
	"fmt"
)

// Type World is a map of City objects
type World map[string]*City

// Add a city to the World
func (world World) AddCityToWorld(city City) *City {
	// City name key points to the passed city object
	world[city.CityName] = &city
	return &city
}

// AddNewCityToWorld takes a named city string and adds it to world map
func (world World) AddNewCityToWorld(name string) *City {
	return world.AddCityToWorld(CreateCity(name))
}

// World as a string
func (world World) String() string {
	var worldString string
	for _, city := range world {
		worldString += fmt.Sprintf("%s\n", city)
	}
	return worldString
}
