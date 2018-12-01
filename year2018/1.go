package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

var seen [2000000]bool
var firstRepeatFound = false
var firstRepeat = 0
var frequency = 0

// DayOne is the solution for December 1st, 2018
func DayOne(input string) string {
	changeStrings := strings.Split(input, "\n")

	// Calibrate once to find the resulting frequency
	calibrate(changeStrings)
	finalFrequency := frequency

	// For part two, we need to keep calibrating until we find a repeat
	for !firstRepeatFound {
		calibrate(changeStrings)
	}

	return fmt.Sprintf("%d (first repeat at %d)", finalFrequency, firstRepeat)
}

func calibrate(changeStrings []string) {
	// Iterate over each change in the list
	for _, changeString := range changeStrings {
		// Convert the change to an int and update the frequency
		change, _ := strconv.Atoi(changeString)
		frequency += change

		// If we've seen this frequency before, and haven't already found a repeat
		if !firstRepeatFound && seen[frequency+1000000] {
			// Record this first repeat
			firstRepeatFound = true
			firstRepeat = frequency
		} else {
			// Record the fact that we've seen a new frequency
			seen[frequency+1000000] = true
		}
	}
}
