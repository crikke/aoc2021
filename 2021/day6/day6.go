package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part1: %d \n", part1(80, input))
	fmt.Printf("part1: %d \n", part1(256, input))
}

func part1(duration int, input string) int {
	strs := strings.Split(input, ",")
	var fishes [9]int

	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		fishes[num]++
	}

	for day := 0; day < duration; day++ {
		fishes = stepDay(fishes)
	}

	sum := 0

	for _, num := range fishes {
		sum += num
	}

	return sum
}

func stepDay(arr [9]int) [9]int {
	return [9]int{arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[0] + arr[7], arr[8], arr[0]}
}

var input = `3,4,3,1,2,1,5,1,1,1,1,4,1,2,1,1,2,1,1,1,3,4,4,4,1,3,2,1,3,4,1,1,3,4,2,5,5,3,3,3,5,1,4,1,2,3,1,1,1,4,1,4,1,5,3,3,1,4,1,5,1,2,2,1,1,5,5,2,5,1,1,1,1,3,1,4,1,1,1,4,1,1,1,5,2,3,5,3,4,1,1,1,1,1,2,2,1,1,1,1,1,1,5,5,1,3,3,1,2,1,3,1,5,1,1,4,1,1,2,4,1,5,1,1,3,3,3,4,2,4,1,1,5,1,1,1,1,4,4,1,1,1,3,1,1,2,1,3,1,1,1,1,5,3,3,2,2,1,4,3,3,2,1,3,3,1,2,5,1,3,5,2,2,1,1,1,1,5,1,2,1,1,3,5,4,2,3,1,1,1,4,1,3,2,1,5,4,5,1,4,5,1,3,3,5,1,2,1,1,3,3,1,5,3,1,1,1,3,2,5,5,1,1,4,2,1,2,1,1,5,5,1,4,1,1,3,1,5,2,5,3,1,5,2,2,1,1,5,1,5,1,2,1,3,1,1,1,2,3,2,1,4,1,1,1,1,5,4,1,4,5,1,4,3,4,1,1,1,1,2,5,4,1,1,3,1,2,1,1,2,1,1,1,2,1,1,1,1,1,4`
