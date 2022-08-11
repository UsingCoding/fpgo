package either

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEither_Left(t *testing.T) {
	result := creator()

	assert.True(t, result.IsLeft())
	assert.False(t, result.IsRight())

	assert.NotPanics(t, func() {
		result.Left()
	})
	assert.Panics(t, func() {
		result.Right()
	})

	var mappedLeft bool
	result.MapLeft(func(s string) {
		mappedLeft = true
	})
	result.MapRight(func(e int) {
		assert.Fail(t, "either mapped right")
	})
	assert.True(t, mappedLeft)
}

func TestEither_Right(t *testing.T) {
	type customStruct struct {
		a string
	}

	result := NewRight[int, customStruct](customStruct{a: "42"})

	assert.True(t, result.IsRight())
	assert.False(t, result.IsLeft())

	assert.NotPanics(t, func() {
		result.Right()
	})
	assert.Panics(t, func() {
		result.Left()
	})

	var mappedRight bool
	result.MapRight(func(c customStruct) {
		mappedRight = true
	})
	result.MapLeft(func(i int) {
		assert.Fail(t, "either mapped left")
	})
	assert.True(t, mappedRight)
}

func creator() Either[string, int] {
	return NewLeft[string, int]("42")
}
