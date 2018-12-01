package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

var seen [1000000]bool
var firstRepeatFound = false
var firstRepeat = 0
var frequency = 0

// DayOne is the solution for December 1st, 2018
func DayOne(input string) string {
	changes := parseInput(strings.Split(input, "\n"))

	// Calibrate once to find the resulting frequency
	calibrate(changes)
	finalFrequency := frequency

	// For part two, we need to keep calibrating until we find a repeat
	for !firstRepeatFound {
		calibrate(changes)
	}

	return fmt.Sprintf("%d (first repeat at %d)", finalFrequency, firstRepeat)
}

func parseInput(changeStrings []string) []int {
	// Create a new slice that we can use to store the results
	changes := make([]int, len(changeStrings))

	// Convert each string in the input into an int and add to destination slice
	for index, changeString := range changeStrings {
		changes[index], _ = strconv.Atoi(changeString)
	}

	return changes
}

func calibrate(changeStrings []int) {
	// Iterate over each change in the list
	for _, change := range changeStrings {
		// Convert the change to an int and update the frequency
		frequency += change

		// If we've seen this frequency before, and haven't already found a repeat
		if !firstRepeatFound && seen[frequency+500000] {
			// Record this first repeat
			firstRepeatFound = true
			firstRepeat = frequency
		} else {
			// Record the fact that we've seen a new frequency
			seen[frequency+500000] = true
		}
	}
}
