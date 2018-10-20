package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/oliverdunk/adventofcode-go/utils"
	"github.com/oliverdunk/adventofcode-go/year2017"
)

var puzzles = make(map[int]map[int]func(string) string)

func main() {
	fmt.Println("🎉 We're up and running!")

	year, day := getDesiredPuzzle()

	fmt.Println()
	fmt.Print("🔵 Attempting to run ", year, " day ", day, ".\n")
	loadPuzzles()
	runPuzzle(year, day)
}

func getDesiredPuzzle() (int, int) {
	if len(os.Args) > 1 && strings.Compare(os.Args[1], "-today") == 0 {
		date := time.Now()

		if date.Month() == 12 {
			return date.Year(), date.Day()
		}

		// We didn't return, so it must be a different month
		fmt.Println("⚠️  Ignoring today switch, which only works in December.")
	}

	year := utils.InputNumber("Please enter the year you're interested in")
	day := utils.InputNumber("Ok, " + strconv.Itoa(year) + " it is. Enter the day")
	return year, day
}

func loadPuzzles() {
	puzzles[2017] = make(map[int]func(string) string)
	puzzles[2017][6] = year2017.DaySix
}

func runPuzzle(year int, day int) {
	if puzzles[year] == nil || puzzles[year][day] == nil {
		fmt.Println("🔴 Unable to find puzzle.")
		return
	}

	filePath := "year" + strconv.Itoa(year) + "/inputs/" + strconv.Itoa(day) + ".txt"
	puzzleInput, error := utils.ReadFile(filePath)

	if error != nil {
		puzzleInput = utils.InputString("Please enter the puzzle input")
		utils.SaveFile(filePath, puzzleInput)
	} else {
		fmt.Println("🔵 Using input from file.")
	}

	result := puzzles[year][day](puzzleInput)
	fmt.Println("🏁 Result:", result)
}
