package year2018

import (
	"fmt"
	"strings"
)

// DayTwo is the solution for December 2nd, 2018
func DayTwo(input string) string {
	checksum := checksum(input)
	commonLetters := findCorrectID(input)

	return fmt.Sprintf("%s (Checksum %d)", commonLetters, checksum)
}

func checksum(input string) int {
	// Store the number of ids encountered with the given number of repeats
	containsLetterTwiceCount := 0
	containsLetterThreeTimesCount := 0

	// Iterate over each id
	for _, id := range strings.Split(input, "\n") {
		// Characters mapped to index by ASCII code - 97
		var appearances [26]int

		// Iterate over each character and increment its appearance count
		for index := 0; index < len(id); index++ {
			appearances[id[index]-97]++
		}

		foundCharWithTwoAppearances := false
		foundCharWithThreeAppearances := false

		for index := 0; index < 26; index++ {
			if appearances[index] == 2 && !foundCharWithTwoAppearances {
				foundCharWithTwoAppearances = true
				containsLetterTwiceCount++
			}
			if appearances[index] == 3 && !foundCharWithThreeAppearances {
				foundCharWithThreeAppearances = true
				containsLetterThreeTimesCount++
			}
		}
	}

	return containsLetterTwiceCount * containsLetterThreeTimesCount
}

// findCorrectId returns the letters two correct box IDs have in common
func findCorrectID(input string) string {
	ids := strings.Split(input, "\n")

	// Iterate over each ID
	for index, id := range ids {
		// See if any subsequent IDs are similar to this one
		for futureIndex := index + 1; futureIndex < len(ids); futureIndex++ {
			tooDifferent := false
			firstDifferenceIndex := -1

			for charIndex := 0; charIndex < len(id); charIndex++ {
				if id[charIndex] != ids[futureIndex][charIndex] {
					if firstDifferenceIndex == -1 {
						firstDifferenceIndex = charIndex
					} else {
						tooDifferent = true
					}
				}
			}

			if !tooDifferent {
				return strings.Replace(id, string(id[firstDifferenceIndex]), "", 1)
			}
		}
	}
	return ""
}
