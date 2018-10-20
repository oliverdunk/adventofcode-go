package year2017

import (
	"fmt"
	"strconv"
	"strings"
)

// DaySix is the solution for December 6th, 2017
func DaySix(input string) string {
	var banks [16]int

	// Load banks with tab separated input
	splitInput := strings.Split(input, "	")

	for i := 0; i < 16; i++ {
		bankValue, _ := strconv.Atoi(splitInput[i])
		banks[i] = bankValue
	}

	return fmt.Sprint(banks)
}
