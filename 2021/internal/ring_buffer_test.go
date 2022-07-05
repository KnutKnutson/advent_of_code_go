package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RingBuffer(t *testing.T) {
	t.Run("adds and returns 1 value", func(t *testing.T) {
		b := NewRingBuffer[int](3)

		b.Add(1)

		assert.Equal(t, 1, b.Values()[0])
	})

	t.Run("stores full size", func(t *testing.T) {
		b := NewRingBuffer[int](3)

		for i := 1; i < 4; i++ {
			b.Add(i)
		}

		assert.ElementsMatch(t, b.Values(), []int{1, 2, 3})
	})

	t.Run("circles ring, returns in", func(t *testing.T) {
		b := NewRingBuffer[int](3)

		for i := 1; i < 6; i++ {
			b.Add(i)
		}

		assert.ElementsMatch(t, b.Values(), []int{3, 4, 5})
	})
}
