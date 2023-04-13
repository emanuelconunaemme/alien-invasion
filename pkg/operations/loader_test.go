package operations

import (
	. "saga/aliens/pkg/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCityMapSuccess(t *testing.T) {

	cities, err := LoadCities("testdata/correct_cities.txt")
	assert.Nil(t, err, "error reading cities: %v", err)

	foo := &City{Name: "Foo"}
	bar := &City{Name: "Bar"}
	bee := &City{Name: "Bee"}
	quUx := &City{Name: "Qu-ux"}
	baz := &City{Name: "Baz"}

	foo.Neighbors = []*Neighbor{
		{City: bar, Direction: North},
		{City: baz, Direction: West},
		{City: quUx, Direction: South},
	}
	bar.Neighbors = []*Neighbor{
		{City: foo, Direction: South},
		{City: bee, Direction: West},
	}
	bee.Neighbors = []*Neighbor{
		{City: bar, Direction: East},
	}
	quUx.Neighbors = []*Neighbor{
		{City: foo, Direction: North},
	}
	baz.Neighbors = []*Neighbor{
		{City: foo, Direction: East},
	}

	expectedCities := []*City{
		foo,
		bar,
		bee,
		quUx,
		baz,
	}

	assert.Equal(t, len(expectedCities), len(cities))
	for i, city := range cities {
		expectedCity := expectedCities[i]
		assert.Equal(t, expectedCity.Name, city.Name)
		assert.Equal(t, len(expectedCity.Aliens), len(city.Aliens))
		assert.Equal(t, len(expectedCity.Neighbors), len(city.Neighbors))
		for j, neighbor := range city.Neighbors {
			expectedNeighbor := expectedCity.Neighbors[j]
			assert.Equal(t, expectedNeighbor.Direction, neighbor.Direction)
			assert.Equal(t, expectedNeighbor.City.Name, neighbor.City.Name)
		}
	}
}

func TestLoadCityMapFailure(t *testing.T) {
	cities, err := LoadCities("testdata/wrong_cities_1.txt")
	assert.Nil(t, cities)
	assert.NotNil(t, err)
	cities, err = LoadCities("testdata/wrong_cities_2.txt")
	assert.Nil(t, cities)
	assert.NotNil(t, err)
}
