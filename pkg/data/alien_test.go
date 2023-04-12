package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAliens(t *testing.T) {
	n := 10
	aliens := GenerateAliens(n)

	assert.Equal(t, len(aliens), n, "Wrong number of aliens")
	for i := 0; i < n; i++ {
		assert.NotEmpty(t, aliens[i].Name, "Alien name should not be empty")
	}
}
