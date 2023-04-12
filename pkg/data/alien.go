package data

import (
	"time"

	"github.com/goombaio/namegenerator"
)

type Alien struct {
	Name string
}

func GenerateAliens(n int) []*Alien {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	aliens := make([]*Alien, n)
	for i := 0; i < n; i++ {
		aliens[i] = &Alien{Name: nameGenerator.Generate()}
	}
	return aliens
}
