# Advent of Code
This contains my [Advent of Code](https://adventofcode.com/) solutions for 2018, which I chose to write in Go.

For solutions to previous years, see my [old repository](https://github.com/oliverdunk/adventofcode) containing Java.

## Running

Clone this repository:

```bash
go get github.com/oliverdunk/adventofcode-go
```

Navigate to the `adventofcode` subdirectory, which contains the main command for running the puzzles:

```bash
cd $GOPATH/src/github.com/oliverdunk/adventofcode-go/adventofcode
```

Build the program:

```bash
go build .
```

Run the program:

```bash
./adventofcode
```

### Flags

The `-today` flag can be used to automatically select today's puzzle. When omitted, the program will prompt you for the year and day of the puzzle you wish to run.

### Inputs

When a puzzle is run for the first time, you will be prompted for your input. This will then be saved in the `yearYYYY/inputs/D.txt` file, where it will be loaded from in the future.

## Adding a puzzle

The entry point for each solution is the `yearYYYY/D.go` file. The main function should take a string as an input and return a string as its result. It should be named `Day` followed by the puzzle day in words.

To register this solution with the main command, edit `adventofcode/main.go`. In particular, add your puzzle to the map populated in the `loadPuzzles` function.

## Contributing

This repository is for my personal solutions, meaning puzzle solutions are unlikely to be accepted. Bug fixes, style changes and other similar PRs are more than welcome.