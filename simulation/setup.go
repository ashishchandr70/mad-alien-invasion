package simulation

import (
	"bufio"
	"fmt"
	"mad-alien-invasion/simulation/types"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// LoadedWorldMap is a list of cities as read from the input file
type LoadedWorldMap []*types.City

// Get remaining Cities from LoadedWorldMap
func (in LoadedWorldMap) GetRemainingCities(world types.World) LoadedWorldMap {
	out := make(LoadedWorldMap, 0, len(in))
	processed := make(map[string]bool)
	for _, city := range in {

		if processed[city.CityName] {
			continue
		}
		if !city.CheckCityDestroyed() {
			out = append(out, city)
			processed[city.CityName] = true
			continue
		}

		if city.East != nil && !city.East.CheckCityDestroyed() && !processed[city.East.CityName] {
			out = append(out, city.East)
			processed[city.East.CityName] = true
		}
		if city.West != nil && !city.West.CheckCityDestroyed() && !processed[city.West.CityName] {
			out = append(out, city.West)
			processed[city.West.CityName] = true
		}
		if city.North != nil && !city.North.CheckCityDestroyed() && !processed[city.North.CityName] {
			out = append(out, city.North)
			processed[city.North.CityName] = true
		}
		if city.South != nil && !city.South.CheckCityDestroyed() && !processed[city.South.CityName] {
			out = append(out, city.South)
			processed[city.South.CityName] = true
		}
	}
	return out
}

// CreateAliensFromRandomSeed creates N new Aliens with random names
func CreateAliensFromRandomSeed(n int, r *rand.Rand) []*types.Alien {
	aliens := make([]*types.Alien, n)
	for i := 0; i < n; i++ {
		name := fmt.Sprintf("Alien-%s", strconv.Itoa(r.Int()))
		alien := types.CreateNewAlien(name)
		aliens[i] = &alien
	}
	return aliens
}

// helper function to return file contents as an array of strings
func scanFileByLine(filename string) ([]string, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("scanFileByLine: could not find file %s. Check to ensure it exists", filename)
	}
	readFile, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("scanFileByLine: could not open file %s. Check to ensure it exists and is readable", filename)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines, nil
}

// Name aliens from a given file
func NameAliens(aliens []*types.Alien, file string) error {
	alienNames, err := scanFileByLine(file)
	if err != nil {
		return err
	}

	for i := 0; i < len(aliens) && i < len(alienNames); i++ {
		aliens[i].Name = alienNames[i]
	}
	return nil
}

// LoadWorldMapFromFile loads the world map from file
func LoadWorldMapFromFile(file string) (types.World, LoadedWorldMap, error) {
	cityNames, err := scanFileByLine(file)
	if err != nil {
		return nil, nil, err
	}

	w := make(types.World)
	input := make(LoadedWorldMap, 0)

	for _, wmap := range cityNames {
		data := strings.Split(wmap, " ")

		city := w.AddNewCityToWorld(data[0])

		for _, s := range data[1:] {
			roadName, cityName, err := parseCityString(s)
			if err != nil {
				return nil, nil, err
			}

			other, exists := w[cityName]
			if !exists {
				other = w.AddNewCityToWorld(cityName)
			}

			// adding our node pointers to other cities
			switch roadName {
			case "north":
				city.North = other
				other.South = city
			case "south":
				city.South = other
				other.North = city
			case "east":
				city.East = other
				other.West = city
			case "west":
				city.West = other
				other.East = city
			}
		}
		input = append(input, city)

	}

	return w, input, nil
}

// Parses a string of cities and connected cities
func parseCityString(input string) (string, string, error) {

	parsedString := strings.SplitN(input, "=", 2)

	return parsedString[0], parsedString[1], nil
}

// Print the world map file
func (in LoadedWorldMap) String() string {
	var out string
	for _, city := range in {
		if !city.CheckCityDestroyed() {
			out += fmt.Sprintf("%s\n", city)
		}

	}
	return out
}
