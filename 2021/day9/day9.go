package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// lows := day9("input")
	// fmt.Printf("part1: %d \n", sum1)
	// fmt.Printf("part2: %d \n", sum2)
}

func day9(input string) (lows int) {

	grid := makeGrid(input)

	for y := 0; y < len(grid); y++ {

		for x := 0; x < len(grid[y]); x++ {

			cell := grid[y][x]

			if cell == nil {
				ninenine := 9999999999
				grid[y][x] = &ninenine
				continue
			}
		}
	}

	for y := 1; y < len(grid)-1; y++ {

		for x := 1; x < len(grid[y])-1; x++ {

			cell := *grid[y][x]

			if cell < *grid[y-1][x] &&
				cell < *grid[y+1][x] &&
				cell < *grid[y][x-1] &&
				cell < *grid[y][x+1] {
				lows += cell + 1
				fmt.Println(x, y, cell)
			}
		}
	}

	return
}

func makeGrid(input string) [][]*int {
	rows := strings.Split(input, "\n")
	grid := make([][]*int, len(rows)+2)

	l := 0
	for y, row := range rows {

		l = len(row) + 2
		grid[y+1] = make([]*int, len(row)+2)
		for x, c := range strings.Split(row, "") {

			val, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			grid[y+1][x+1] = &val

		}
	}

	grid[0] = make([]*int, l)
	grid[len(rows)+1] = make([]*int, l)

	return grid
}
