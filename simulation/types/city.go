package types

import (
	"fmt"
)

// City has a name, a state and points to other cities via City pointers
type City struct {
	CityName        string
	IsCityDestroyed bool
	East            *City
	West            *City
	North           *City
	South           *City
	InvaldingAliens []*Alien
}

// CreateCity creates a City with a name and default state of not destroyed, and nil pointers to other cities
func CreateCity(name string) City {
	return City{
		CityName:        name,
		InvaldingAliens: make([]*Alien, 0),
	}
}

// CheckCityDestroyed checks if City is destroyed
func (c *City) CheckCityDestroyed() bool {
	return c.IsCityDestroyed
}

// DestroyCity sets the IsCityDestroyed flag indicating city has been destroyed
func (c *City) DestroyCity() {
	c.IsCityDestroyed = true
}

// CityInvaded sets the counter for number of aliens invading the city
func (c *City) CityInvaded(a *Alien) {
	c.InvaldingAliens = append(c.InvaldingAliens, a)
}

// CityAbandoned decrements the counter for number of aliens invading the city
func (c *City) CityAbandoned(a *Alien) {
	for index, alien := range c.InvaldingAliens {
		if a == alien {
			c.InvaldingAliens = append(c.InvaldingAliens[:index], c.InvaldingAliens[index+1:]...)
		}
	}
}

// String printing of a City
func (c *City) String() string {
	var links string
	if c.East != nil && !c.East.CheckCityDestroyed() {
		links += fmt.Sprintf("east=%s ", c.East.CityName)
	}
	if c.West != nil && !c.West.CheckCityDestroyed() {
		links += fmt.Sprintf("west=%s ", c.West.CityName)
	}
	if c.South != nil && !c.South.CheckCityDestroyed() {
		links += fmt.Sprintf("south=%s ", c.South.CityName)
	}
	if c.North != nil && !c.North.CheckCityDestroyed() {
		links += fmt.Sprintf("north=%s ", c.North.CityName)
	}
	if len(links) == 0 {
		return c.CityName
	}
	return fmt.Sprintf("%s %s", c.CityName, links[:len(links)-1])
}
