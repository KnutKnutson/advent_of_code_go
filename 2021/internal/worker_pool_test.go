package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkerPool(t *testing.T) {
	t.Run("queue full returns error", func(t *testing.T) {
		nothing := func() {}
		pool := NewWorkerPool(1, 2)

		// fill the queue of size 2
		require.Nil(t, pool.Do(nothing))
		require.Nil(t, pool.Do(nothing))

		// adding third is error
		require.NotNil(t, pool.Do(nothing))
	})
}
