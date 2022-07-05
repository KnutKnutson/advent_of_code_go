package main

import (
	"testing"
	time2 "time"

	"github.com/stretchr/testify/assert"
)

func Test_depthAnalyzer(t *testing.T) {
	t.Run("analyzeDepth", func(t *testing.T) {
		t.Run("tracks depth changes correctly", func(t *testing.T) {
			da := newDepthAnalyzer()
			for _, depth := range []int{1, 5, 4, 4, 6} {
				da.analyzeDepth(depth)
			}
			assert.Equal(t, da.getNumIncreases(), 2)
		})

		t.Run("tracks the sliding window depth changes correctly", func(t *testing.T) {
			da := newDepthAnalyzer()
			for _, depth := range []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263} {
				da.analyzeDepth(depth)
			}
			assert.Equal(t, da.getNumIncreasesSliding(), 5)
		})

		t.Run("tracks the sliding window depth changes correctly", func(t *testing.T) {
			time := time2.Now()
			unix := time.Unix() * 1000
			unixMilli := time.UnixMilli()
			assert.Equal(t, unix, unixMilli)
		})
	})
}
