## Mad Alien Invasion

  

Mad aliens are about to invade the earth.


## Design Basics
The following are the types of objects/entities used in the simulation.

### Alien

- Has a name
- Has a state i.e. alive or dead
- Optionally invades a city

### City

- Has a name
- Has a state i.e. destroyed or not
- Points to other cities in north, south, east or west, if those cities exist
- Is invaded by 1 to 2 aliens

### World

- Is a map of City objects with a key being the name of the city. This disallows duplicate cities being mapped into the world


### Prerequisite

Setup [Golang](https://golang.org/doc/install)
  
  
### Usage
Clone the repository

```bash
git clone https://github.com/ashishchandr70/mad-alien-invasion.git
cd mad-alien-invasion
```

  

### Run
Getting started
```
$ go run main.go -aliens 20
```
**Options**: View all CLI options
  
```bash
$ go run main.go -help
```

## Testing
Run the test suite with:

```
$ go test ./...
```

## Assumptions

1. Only 2 aliens can invade a city at a time, which results in the destruction of the city and the aliens themselves getting killed.


