package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"knutknut.com/advent_of_code_go/2021/internal"
)

type depthAnalyzer struct {
	numIncreases, numIncreasesSliding int
	lastDepth, lastSlidingDepth       int
	slidingWindow                     *internal.RingBuffer[int]
}

func newDepthAnalyzer() *depthAnalyzer {
	return &depthAnalyzer{
		lastDepth:        -1,
		lastSlidingDepth: -1,
		slidingWindow:    internal.NewRingBuffer[int](3),
	}
}

func (d *depthAnalyzer) analyzeDepth(depth int) {
	if d.lastDepth > -1 && d.lastDepth < depth {
		d.numIncreases++
	}
	d.lastDepth = depth

	d.slidingWindow.Add(depth)
	if d.slidingWindow.CurrentSize() == 3 {
		slidingDepth := 0
		for _, v := range d.slidingWindow.Values() {
			slidingDepth += v
		}
		if d.lastSlidingDepth > -1 && d.lastSlidingDepth < slidingDepth {
			d.numIncreasesSliding++
		}
		d.lastSlidingDepth = slidingDepth
	}
}

func (d *depthAnalyzer) getNumIncreases() int {
	return d.numIncreases
}

// part 2 of day 1. analyze depth with a sliding window
func (d *depthAnalyzer) getNumIncreasesSliding() int {
	return d.numIncreasesSliding
}

// count the number of times a depth measurement increases from the previous measurement.
// (There is no measurement before the first measurement.)
// How many measurements are larger than the previous measurement?
func main() {
	file, err := os.Open("part1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	da := newDepthAnalyzer()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		da.analyzeDepth(depth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Num Depth Increases: %v", da.getNumIncreases())
	log.Printf("Num Sliding Depth Increases: %v", da.getNumIncreasesSliding())
}
