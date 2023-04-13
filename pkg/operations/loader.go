package operations

import (
	"bufio"
	"fmt"
	"os"
	. "saga/aliens/pkg/data"
)

func LoadCities(filePath string) ([]*City, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// loading line by line initializes cities with sparse neighbors (not a graph)
	cities := []*City{}

	for scanner.Scan() {
		line := scanner.Text()
		city, err := NewCityFromString(line)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// hidrate neightbors with real references
	for _, city := range cities {
		for _, neighbor := range city.Neighbors {
			hydratedCity := getCity(cities, neighbor.City.Name)
			if hydratedCity == nil {
				return nil, fmt.Errorf("Can't find %s in the city list", neighbor.City.Name)
			}
			neighbor.City = hydratedCity
		}
	}

	if err := validateData(cities); err != nil {
		return nil, err
	}

	return cities, nil
}

func getCity(cities []*City, cityName string) *City {
	for _, city := range cities {
		if city.Name == cityName {
			return city
		}
	}
	return nil
}

func validateData(cities []*City) error {
	for _, city := range cities {
		for _, neighbor := range city.Neighbors {

			neighborCity := neighbor.City
			neighborDirection := neighbor.Direction

			otherNeighbor := neighborCity.GetNeighbor(neighborDirection.Opposite())
			if otherNeighbor == nil {
				return fmt.Errorf("%s => %s, but NOT %s => %s", city.Name, neighbor.City.Name, neighbor.City.Name, city.Name)
			}

			otherCity := neighborCity.GetNeighbor(neighborDirection.Opposite()).City

			if city.Name != otherCity.Name {
				// When A is noth of B, B has to be south of A
				return fmt.Errorf("%s => %s, but NOT %s => %s", city.Name, neighbor.City.Name, neighbor.City.Name, city.Name)
			}
		}
	}
	return nil
}
