package data

import (
	"errors"
	"fmt"
	"strings"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

var directionsMap = map[string]Direction{
	"north": North,
	"east":  East,
	"south": South,
	"west":  West,
}

// redundant data, but given the scope, this is a tradeoff for efficiency
var directionsStringMap = map[Direction]string{
	North: "north",
	East:  "east",
	South: "south",
	West:  "west",
}

type City struct {
	Name      string
	Neighbors []*Neighbor
	Aliens    []*Alien
}

type Neighbor struct {
	City      *City
	Direction Direction
}

func NewCityFromString(input string) (*City, error) {
	// Split input string by space to get city name and its neighbors
	fields := strings.Fields(input)
	if len(fields) < 2 {
		return nil, errors.New("invalid input string")
	}
	cityName := fields[0]
	neighborFields := fields[1:]

	// Parse neighbors
	neighbors := []*Neighbor{}
	for _, neighborField := range neighborFields {
		parts := strings.Split(neighborField, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid neighbor field: %s", neighborField)
		}
		directionStr := parts[0]
		direction, ok := directionsMap[directionStr]
		if !ok {
			return nil, fmt.Errorf("invalid direction: %s", directionStr)
		}
		neighbors = append(neighbors, &Neighbor{
			City:      &City{Name: parts[1]},
			Direction: direction,
		})
	}

	return &City{
		Name:      cityName,
		Neighbors: neighbors,
	}, nil
}

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("Invalid direction")
	}
}

func (c *City) HasFight() bool {
	return c.Aliens != nil && len(c.Aliens) > 1
}

// pardon my Kotlin/Java and there is a more idiomatic way to do it
func (d *Direction) ToString() string {
	return directionsStringMap[*d]
}

func (c *City) ToString() string {
	neighborsStr := ""
	for _, neighbor := range c.Neighbors {
		neighborsStr += fmt.Sprintf(" %s=%s", neighbor.Direction.ToString(), neighbor.City.Name)
	}
	return fmt.Sprintf("%s%s", c.Name, neighborsStr)
}

func (c *City) HasAlien(alien *Alien) bool {
	if c.Aliens == nil {
		return false
	}
	for _, a := range c.Aliens {
		if alien.Name == a.Name {
			return true
		}
	}
	return false
}

func (c *City) RemoveAlien(alien *Alien) bool {
	for i, a := range c.Aliens {
		if alien.Name == a.Name {
			c.Aliens = append(c.Aliens[:i], c.Aliens[i+1:]...)
			return true
		}
	}
	return false
}

func (c *City) AddAlien(alien *Alien) {
	if c.Aliens == nil {
		c.Aliens = []*Alien{}
	}
	c.Aliens = append(c.Aliens, alien)
}
