package types

import "fmt"

/*
Alien has a name, a state (dead or alive) and optionally occupies a city
*/
type Alien struct {
	Name        string
	IsAlienDead bool
	city        *City
}

// CreateNewAlien creates an Alien with a name and default flags
func CreateNewAlien(name string) Alien {
	return Alien{
		Name: name,
	}
}

// InvadeCity changes the alien's city to the invaded city
func (a *Alien) InvadeCity(city *City) {
	a.city = city
	city.CityInvaded(a)
}

// LeaveCity decrements the alien count invading the city
func (a *Alien) LeaveCity(city *City) {
	city.CityAbandoned(a)
}

// City this Alien is occupying
func (a *Alien) City() *City {
	return a.city
}

// IsDead returns alien's state - dead or alive
func (a *Alien) IsDead() bool {
	return a.IsAlienDead
}

// Kill sets the alien's state to dead
func (a *Alien) Kill() {
	a.IsAlienDead = true
}

// IsCityUnderInvasion checks if Alien is currently invading a City
func (a *Alien) IsCityUnderInvasion() bool {
	return a.city != nil
}

// IsTrapped checks if Alien is trapped
func (a *Alien) IsTrapped() bool {
	if !a.IsCityUnderInvasion() {
		return false
	}

	if a.city.East != nil && !a.city.East.CheckCityDestroyed() {
		return false
	} else if a.city.West != nil && !a.city.West.CheckCityDestroyed() {
		return false
	} else if a.city.South != nil && !a.city.South.CheckCityDestroyed() {
		return false
	} else if a.city.North != nil && !a.city.North.CheckCityDestroyed() {
		return false
	}

	return true
}

// String function prints an alien
func (a *Alien) String() string {
	return fmt.Sprintf("name=%s city={%s}\n", a.Name, a.city)
}
