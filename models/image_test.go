package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageStorage_GetSizes(t *testing.T) {
	c := NewConfig()
	sizes := ImageStorage{}.GetSizes()
	assert.Equal(t, c.ImageSmallWidth*2, sizes["small"].Width)
	assert.Equal(t, c.ImageSmallHeight*2, sizes["small"].Height)
	assert.Equal(t, c.ImageSmallWidth, sizes["small@ld"].Width)
	assert.Equal(t, c.ImageSmallHeight, sizes["small@ld"].Height)

	assert.Equal(t, c.ImageMiddleWidth*2, sizes["middle"].Width)
	assert.Equal(t, c.ImageMiddleHeight*2, sizes["middle"].Height)
	assert.Equal(t, c.ImageMiddleWidth, sizes["middle@ld"].Width)
	assert.Equal(t, c.ImageMiddleHeight, sizes["middle@ld"].Height)

	assert.Equal(t, c.ImageBigWidth*2, sizes["big"].Width)
	assert.Equal(t, c.ImageBigHeight*2, sizes["big"].Height)
	assert.Equal(t, c.ImageBigWidth, sizes["big@ld"].Width)
	assert.Equal(t, c.ImageBigHeight, sizes["big@ld"].Height)
}
