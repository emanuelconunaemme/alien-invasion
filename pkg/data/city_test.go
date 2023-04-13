package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCitySuccess(t *testing.T) {
	input := "Foo north=Bar west=Baz south=Qu-ux"
	expectedCity := &City{
		Name: "Foo",
		Neighbors: []*Neighbor{
			{
				City:      &City{Name: "Bar"},
				Direction: North,
			},
			{
				City:      &City{Name: "Baz"},
				Direction: West,
			},
			{
				City:      &City{Name: "Qu-ux"},
				Direction: South,
			},
		},
	}
	city, err := NewCityFromString(input)
	assert.Nil(t, err, "Unexpected error")
	assert.EqualValues(t, city, expectedCity)
}

func TestInitCityMalformed(t *testing.T) {
	assertNewCityError(t, "Foo Bar=north")
	assertNewCityError(t, "Foo")
	assertNewCityError(t, "Foo north=Bar wrongdirection=Baz")
}

func assertNewCityError(t *testing.T, input string) {
	city, err := NewCityFromString(input)
	assert.Nil(t, city, "expecting an error")
	assert.NotNil(t, err, "expecting an error")
}

func TestCityToString(t *testing.T) {
	input := "Foo north=Bar west=Baz south=Qu-ux"
	city, err := NewCityFromString(input)
	assert.Nil(t, err, "unexpected error")
	cityString := city.ToString()
	assert.Equal(t, cityString, input, "Wrong output")
}

func TestHasFight(t *testing.T) {
	input := "Foo north=Bar west=Baz south=Qu-ux"
	city, err := NewCityFromString(input)
	assert.Nil(t, err, "Unexpected error")
	assert.NotNil(t, city, "City not initialized")
	assert.False(t, city.HasFight(), "City should have no fights")

	alien1 := &Alien{Name: "Alien1"}
	alien2 := &Alien{Name: "Alien2"}
	city.Aliens = []*Alien{alien1}
	assert.False(t, city.HasFight(), "City should have no fights")
	city.Aliens = []*Alien{alien1, alien2}
	assert.True(t, city.HasFight(), "City should have fights")
}

func TestAliens(t *testing.T) {
	input := "Foo north=Bar west=Baz south=Qu-ux"
	city, _ := NewCityFromString(input)
	alien1 := &Alien{Name: "Alien1"}
	alien2 := &Alien{Name: "Alien2"}

	assert.False(t, city.HasAlien(alien1), "city shouldn't have aliens")
	assert.False(t, city.HasAlien(alien2), "city shouldn't have aliens")

	city.AddAlien(alien1)
	assert.True(t, city.HasAlien(alien1), "city should have alien1")
	assert.False(t, city.HasAlien(alien2), "city shouldn't have alien1")

	city.AddAlien(alien2)
	assert.True(t, city.HasAlien(alien1), "city should have alien1")
	assert.True(t, city.HasAlien(alien2), "city should have alien2")

	city.RemoveAlien(alien1)
	assert.False(t, city.HasAlien(alien1), "city shouldn't have alien1")
	assert.True(t, city.HasAlien(alien2), "city should have alien2")
}

func TestConnection(t *testing.T) {
	// isolated
	city := &City{
		Name: "Foo",
	}
	assert.True(t, city.IsIsolated())
	assert.False(t, city.IsConnected())

	// connected
	city = &City{
		Name: "Foo",
		Neighbors: []*Neighbor{
			{
				City:      &City{Name: "Bar"},
				Direction: North,
			},
		},
	}
	assert.False(t, city.IsIsolated())
	assert.True(t, city.IsConnected())
}
