package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/oliverdunk/adventofcode-go/utils"
	"github.com/oliverdunk/adventofcode-go/year2018"
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
	todayPtr := flag.Bool("today", false, "If today's puzzle should be selected automatically")
	flag.Parse()

	if *todayPtr {
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
	puzzles[2018] = make(map[int]func(string) string)
	puzzles[2018][1] = year2018.DayOne
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
