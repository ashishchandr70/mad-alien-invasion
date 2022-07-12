package main

import (
	"errors"
	"flag"
	"fmt"
	"mad-alien-invasion/simulation"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Required constants
const nDefaultMovements int = 10000                  // Number of movements if not specified on the command line
const nDefaultAliens int = 0                         // Number of aliens who will fight, if not specified on the command line
const sWorldMapFile string = "./inputs/worldmap.txt" // Name of the world file

// Optional constants
const sAlienNames = "./inputs/aliennames.txt" // Names of the aliens (not required but can be used)

// Don't really need a struct here but this is cleaner if we want to add more command line args
type CommandLineArgs struct {
	numAliens int
}

// global var for handling command line args
var cla CommandLineArgs

func init() {
	flag.IntVar(&cla.numAliens, "aliens", nDefaultAliens, "number of aliens to create for simulation")
	flag.Parse()
}

func validateInput() error {
	// check passed value of aliens as a command line arg
	if cla.numAliens <= 0 {
		return errors.New("number of aliens specified is invalid. It must be a valid number > 0")
	}
	if _, err := os.Stat(sWorldMapFile); err != nil {
		return errors.New("input file for the world map does not exist. Check the file inputs/worldmap.txt and ensure it is a valid input file")
	}
	return nil
}

func simulateAlienInvasion() {
	if err := validateInput(); err != nil {
		fmt.Printf("error validating input: %s\n", err)
		flag.Usage()
		os.Exit(1)
	}

	world, in, err := simulation.LoadWorldMapFromFile(sWorldMapFile)
	if err != nil {
		fmt.Printf("File %s exists but could not be used to create the world map. Check its contents. Exiting", sWorldMapFile)
		os.Exit(1)
	}

	randomGenerator := randomize()
	aliens := simulation.CreateAliensFromRandomSeed(cla.numAliens, randomGenerator)

	// associate alien names to number of aliens, if we have a file of alien names
	if _, err := os.Stat(sAlienNames); err == nil {
		if err := simulation.NameAliens(aliens, sAlienNames); err != nil {
			fmt.Printf("Could not associate aliens with names. Check file %s. Error: %s. This is not a showstopper. Continuing\n", sAlienNames, err)
		}
	}
	sim := simulation.NewSimulation(randomGenerator, nDefaultMovements, world, aliens)
	// Start the alien invasion simulation
	if err := sim.StartSimulation(); err != nil {
		fmt.Printf(formatImportantMessage("Error while running simulation: %s"), err)
		os.Exit(1)
	}
	// Complete
	fmt.Println(formatImportantMessage("Invasion Completed Cities left"))
	fmt.Print(in.GetRemainingCities(world))
}

// random number generator for current timestamp in nanoseconds
func randomize() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// formatImportantMessage formats an important message ;)
func formatImportantMessage(msg string) string {
	line := strings.Repeat("+", len(msg))
	out := fmt.Sprintf("\n")
	out += fmt.Sprintf("%s\n", line)
	out += fmt.Sprintf("%s\n", msg)
	out += fmt.Sprintf("%s\n", line)
	return out
}

func main() {
	simulateAlienInvasion()
}
