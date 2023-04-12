package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerminalState(t *testing.T) {
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
	baz := &City{
		Name: "baz",
	}
	cities := []*City{
		foo,
		bar,
		baz,
	}
	cities[2].Aliens = aliens // two aliens in "Baz"

	// all city destroyed
	state := State{Cities: nil, Aliens: aliens}
	assert.True(t, state.IsTerminal())
	// all aliens destroyed
	state = State{Cities: cities, Aliens: nil}
	assert.True(t, state.IsTerminal())
	// disconnected city (2 aliens in baz cannot go anywhere)
	state = State{Cities: cities, Aliens: aliens}
	assert.True(t, state.IsTerminal())

	// moving aliens to "Foo" which is connected to "Bar"
	cities[2].Aliens = nil
	cities[0].Aliens = aliens
	assert.False(t, state.IsTerminal())
}
