package simulation

import (
	"fmt"
	"mad-alien-invasion/simulation/types"
	"mad-alien-invasion/utils"
	"math/rand"
	"sort"
)

// Collection of all Aliens
type Aliens []*types.Alien

// Simulation struct represents a running invasion game
type Simulation struct {
	R            *rand.Rand
	Iteration    int
	EndIteration int
	types.World
	Aliens
}

// NoOpReason why Alien did not make a move
type NoOpReason uint8

// NoOpError when next move can not be made
type NoOpError struct {
	reason NoOpReason
}

const (
	// NoOpAlienDead when Alien is Dead
	NoOpAlienDead NoOpReason = iota
	// NoOpAlienTrapped when Alien is Trapped
	NoOpAlienTrapped
	// NoOpWorldDestroyed when World is destroyed
	NoOpWorldDestroyed
	// NoOpMessage when no-op
	NoOpMessage = " || NO move! %s\n"
)

// Error string representation
func (err *NoOpError) Error() string {
	return fmt.Sprintf("Simulator no-op with reason: %d", err.reason)
}

// Initialize a new Simulation
func NewSimulation(r *rand.Rand, endIteration int, world types.World, aliens Aliens) Simulation {
	return Simulation{
		R:            r,
		Iteration:    0,
		EndIteration: endIteration,
		World:        world,
		Aliens:       aliens,
	}
}

// Starts the alien invasion simulation
func (s *Simulation) StartSimulation() error {
	for ; s.Iteration < s.EndIteration; s.Iteration++ {
		picks := utils.ShuffleLen(len(s.Aliens), s.R)
		noOpRound := true
		for _, p := range picks {
			if err := s.MoveAlien(s.Aliens[p]); err != nil {
				if _, ok := err.(*NoOpError); ok {
					continue
				}
				return err
			}
			noOpRound = false
		}
		if noOpRound {
			return nil
		}
	}
	// Done with the game!
	return nil
}

// MoveAlien moves the Alien to a city
func (s *Simulation) MoveAlien(alien *types.Alien) error {
	fromCity, toCity, err := s.pickMove(alien)
	if err != nil {
		return err
	}

	alien.InvadeCity(toCity)
	if fromCity != nil {
		alien.LeaveCity(fromCity)
	}

	if len(toCity.InvaldingAliens) > 1 {
		toCity.DestroyCity()
		// Kill Aliens and notify
		out := fmt.Sprintf("City named %s has been destroyed by ", toCity.CityName)
		for _, a := range toCity.InvaldingAliens {
			out += fmt.Sprintf("alien %s and ", a.Name)
			a.Kill()
		}
		out = out[:len(out)-5] + "!\n"
		fmt.Print(out)
	}
	// Done
	return nil
}

// pickMove returns Alien move from City to City
func (s *Simulation) pickMove(alien *types.Alien) (*types.City, *types.City, error) {
	// Check if dead or trapped
	fromCity := alien.City()
	if err := checkAlien(alien); err != nil {
		return fromCity, nil, err
	}
	// At the beginning
	if fromCity == nil {
		toCity := s.pickAnyCity()
		if toCity == nil {
			return fromCity, toCity, &NoOpError{reason: NoOpWorldDestroyed}
		}
		return fromCity, toCity, nil
	}
	// Move to next City
	toCity := s.pickConnectedCity(alien)
	if toCity == nil {
		return fromCity, toCity, &NoOpError{reason: NoOpWorldDestroyed}
	}
	return fromCity, toCity, nil
}

// checkAlien returns NoOpError if Alien dead or trapped
func checkAlien(alien *types.Alien) *NoOpError {
	if alien.IsDead() {
		return &NoOpError{NoOpAlienDead}
	}
	if alien.IsTrapped() {
		return &NoOpError{NoOpAlienTrapped}
	}
	return nil
}

// pickConnectedCity picks a random road to an undestroyed City
func (s *Simulation) pickConnectedCity(alien *types.Alien) *types.City {

	if !alien.IsCityUnderInvasion() {
		return nil
	}

	pickMap := make([]string, 0, 4)

	if alien.City().East != nil && !alien.City().East.CheckCityDestroyed() {
		pickMap = append(pickMap, "east")
	}
	if alien.City().West != nil && !alien.City().West.CheckCityDestroyed() {
		pickMap = append(pickMap, "west")
	}
	if alien.City().South != nil && !alien.City().South.CheckCityDestroyed() {
		pickMap = append(pickMap, "south")
	}
	if alien.City().North != nil && !alien.City().North.CheckCityDestroyed() {
		pickMap = append(pickMap, "north")
	}

	if len(pickMap) > 0 {
		pick := s.R.Intn(len(pickMap))
		// Any undestroyed connected city
		switch pickMap[pick] {
		case "north":
			return alien.City().North
		case "south":
			return alien.City().South
		case "east":
			return alien.City().East
		case "west":
			return alien.City().West
		}
	}

	return nil
}

// pickAnyCity picks any undestroyed City in the World
func (s *Simulation) pickAnyCity() *types.City {

	var keys []string
	for k := range s.World {
		if c := s.World[k]; !c.CheckCityDestroyed() {
			keys = append(keys, k)
		}
	}
	if len(keys) == 0 {
		return nil
	}
	sort.Strings(keys)
	pick := s.R.Intn(len(keys))
	return s.World[keys[pick]]
}
