package operations

import (
	. "saga/aliens/pkg/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	aliens := GenerateAliens(2)
	foo := &City{
		Name: "Foo",
		Neighbors: []*Neighbor{
			{
				City:      &City{Name: "Bar"},
				Direction: North,
			},
		},
	}
	bar := &City{
		Name: "Bar",
		Neighbors: []*Neighbor{
			{
				City:      &City{Name: "Foo"},
				Direction: South,
			},
		},
	}
	cities := []*City{
		foo,
		bar,
	}

	state := InitState(cities, aliens)
	assert.Equal(t, cities, state.Cities)
	assert.Equal(t, aliens, state.Aliens)
	// check each alien is in a city
	actualAliens := []*Alien{}
	for _, city := range state.Cities {
		for _, alien := range city.Aliens {
			actualAliens = append(actualAliens, alien)
		}
	}

	assert.Equal(t, len(aliens), len(actualAliens))
}

func TestDestroy(t *testing.T) {
	aliens := GenerateAliens(2)
	foo := &City{
		Name: "Foo",
		Neighbors: []*Neighbor{
			{
				City:      &City{Name: "Bar"},
				Direction: North,
			},
		},
	}
	bar := &City{
		Name: "Bar",
		Neighbors: []*Neighbor{
			{
				City:      &City{Name: "Foo"},
				Direction: South,
			},
		},
	}
	cities := []*City{
		foo,
		bar,
	}
	// two aliens in the first city
	foo.AddAlien(aliens[0])
	foo.AddAlien(aliens[1])
	state := State{Cities: cities, Aliens: aliens}

	// the city should not be there anymore after being destroyed
	assert.Equal(t, 2, len(state.Cities))
	DestroyCity(&state, foo.Name)
	assert.Equal(t, 1, len(state.Cities))
}
