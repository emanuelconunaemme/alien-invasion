package play

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := Play("data/input.txt", 2, true)
	assert.NotNil(t, err)
	err = Play("data/input.txt", 10, true)
	assert.NotNil(t, err)
	err = Play("data/input.txt", 3, false)
	assert.NotNil(t, err)
}
