package operations

import (
	"math/rand"
	. "saga/aliens/pkg/data"
	"time"
)

// this file has all the state mutation operations, not sure this is the idiomatic way in golang

// create a state and randomly place aliens in cities
func InitState(cities []*City, aliens []*Alien) State {
	state := State{Cities: cities, Aliens: aliens}

	// initialize random positions of the aliens
	rand.Seed(time.Now().UnixNano())
	for _, alien := range state.Aliens {
		cityIndex := rand.Intn(len(state.Cities))
		if state.Cities[cityIndex].Aliens == nil {
			state.Cities[cityIndex].Aliens = []*Alien{}
		}
		state.Cities[cityIndex].Aliens = append(state.Cities[cityIndex].Aliens, alien)
	}

	return state
}

func DestroyCity(state *State, cityName string) bool {
	cityIndex := -1
	for i, city := range state.Cities {
		// if city destroy and destroy aliens
		if city.Name == cityName {
			cityIndex = i
			for _, alien := range city.Aliens {
				state.KillAlien(alien)
			}
		} else {
			// remove city from the list of neighbors
			for neighborIndex, neighbor := range city.Neighbors {
				if neighbor.City.Name == cityName {
					city.Neighbors = append(city.Neighbors[:neighborIndex], city.Neighbors[neighborIndex+1:]...)
					break
				}
			}
		}
	}
	if cityIndex != -1 {
		state.Cities = append(state.Cities[:cityIndex], state.Cities[cityIndex+1:]...)
		return true
	}
	return false
}

// move alien and returns the city
func MoveAlien(state *State, alienName string) *City {
	for _, city := range state.Cities {
		for _, alien := range city.Aliens {
			// found the alien
			if alien.Name == alienName {
				// cannot move anywhere
				if city.IsIsolated() {
					return nil
				}
				// remove alien
				city.RemoveAlien(alien)
				// add to the other city
				newCityIndex := rand.Intn(len(city.Neighbors))
				newCity := city.Neighbors[newCityIndex].City
				newCity.AddAlien(alien)
				return newCity
			}
		}
	}
	return nil
}
