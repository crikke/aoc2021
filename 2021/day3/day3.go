package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	gamma, epsilon := part1(input)
	fmt.Println(gamma * epsilon)
	fmt.Println(gamma)
	fmt.Println(epsilon)

	fmt.Println("part 2")
	oxygen, co2 := part2(input)
	fmt.Println(oxygen * co2)
	fmt.Println(oxygen)
	fmt.Println(co2)
}

func part1(input string) (gamma, epsilon int) {
	num := strings.Split(input, "\n")
	occ := make([]int, len(num[0]))

	for _, str := range num {

		for i, c := range str {
			if c == '1' {
				occ[i]++
			}
		}
	}

	for i := 0; i < len(occ); i++ {
		if occ[i] > len(num)/2 {
			gamma += int(math.Pow(2, float64(len(occ)-i-1)))
			continue
		}
	}

	epsilon = gamma
	epsilon ^= int((math.Pow(2, float64(len(occ)))) - 1)

	return
}

// insering items should be at worst O(n)
// getting min/max should be O(h) where h = tree depth
func part2(in string) (oxygen, co2 int) {
	input := strings.Split(in, "\n")

	bitCount := len(input[0])
	root := node{}
	for _, item := range input {
		i, err := strconv.ParseInt(item, 2, 0)

		if err != nil {
			panic(err)
		}
		root.insert(int(i), bitCount)
	}
	oxygen = root.max()
	co2 = root.min()
	return
}

type node struct {
	left  *node
	right *node
	size  int
	key   int
}

func (n *node) insert(data int, bit int) {
	n.size++

	// bit 0 means that this is a leaf and in that case
	// store data as key
	if bit == 0 {
		n.key = data
		return
	}

	// get bit at position
	b := math.Pow(2, float64(bit)-1)

	// check if bit is 1
	if data&int(b) == int(b) {

		if n.right == nil {
			n.right = &node{}
		}

		n.right.insert(data, bit-1)
		return
	} else {
		if n.left == nil {
			n.left = &node{}
		}
		n.left.insert(data, bit-1)
	}
}

func (n *node) max() int {

	if n.key != 0 {
		return n.key
	}

	if n.left != nil && n.right != nil {
		if n.left.size <= n.right.size {
			return n.right.max()
		} else {
			return n.left.max()
		}
	} else {
		if n.left == nil {
			return n.right.max()
		} else {
			return n.left.max()
		}
	}
}

func (n *node) min() int {

	if n.key != 0 {
		return n.key
	}

	if n.left != nil && n.right != nil {

		if n.left.size > n.right.size {
			return n.right.min()
		} else {
			return n.left.min()
		}
	} else {
		if n.left == nil {
			return n.right.min()
		} else {
			return n.left.min()
		}
	}
}

var input = `011110111110
110111000111
001000001010
000101111101
101011011110
010010010111
111101001100
111101011100
110011101101
001111010011
110000000110
111000010001
011011011000
000111110100
111100101010
001011001100
011100000101
100011110110
111001011101
110101011010
110101101111
011011101001
111001010110
011000101000
100101111100
011101110010
010000110000
111010110100
000001000011
110111000000
101010010110
100000001011
001111110001
001111000011
001101001010
100011001111
001011001101
010001010000
000000000000
011001011011
000101100101
010101101101
000110011001
100100110001
010110011011
110101100101
010101101011
011100000000
011001101000
001000100110
101001010011
111111101101
100010111001
000011011100
010001111100
011011111100
000100010001
101111100110
000000111101
111010110001
000100110101
111010100100
111111001011
001000010010
100100000110
010111010110
111111010110
000011011011
001111001000
000000010101
011011100111
110010111110
110001011111
011101100110
010011111100
010101100011
000110010110
100110001111
010111100000
001110110000
101010010010
010011011001
100000100110
010100011010
100110001010
111110001001
101001000000
100010110000
101100001000
000000110101
100001010011
011101100001
000100111001
011110100001
001111000111
001011001111
110111000101
011001010101
101100111000
011111110010
000000100001
100101011111
110110011100
011001011101
000101010001
100100001000
110100001111
101100010010
010000011001
010111011111
001001101101
101110101001
101111011101
100011010010
101001100010
110010111101
100100010001
100010011000
100010110110
100110101101
101110010011
111111011001
100110100111
011111011100
011011001010
101000011110
010110011110
101000111010
110100001101
111011010001
011111100010
101111000010
111011111011
001000011010
010000001101
001101110101
000110010011
011001110110
110010010001
010001010011
110111100100
011100110011
011000010111
100011101110
101100110110
101000100111
100000000111
011011011010
011111011000
010001111101
000010111100
101100100001
011011111111
010111001110
110100111101
011101111010
100010001001
101110011111
000000000101
110100100000
010100011100
101100111011
110011110101
000101110111
001111100011
001110001001
101001001111
010000011101
100101111101
001101000111
000100100110
010110111110
111110110011
101010101011
101010011011
100011010101
101110010100
110110001100
111001010010
011000111000
110011001101
111100100000
100111110011
110011010000
100000101111
110110101001
001010000010
000010100100
101111101001
011111001001
111001111011
001100001110
111011011101
110101001111
010100001101
000000001111
001111001111
100101101001
011101101101
101100101100
110100101001
100111001001
111111001001
101011000011
010101110000
010110100100
101011000100
100010001111
111001101010
111010010010
000011000100
110000111101
110011110111
101111011000
100110110110
000011111010
011011110111
011010010001
110010111000
100101000101
110111011101
110010011110
010101010100
110010011000
101001101100
100111011101
001001100101
011001010001
001101000000
011110111010
001010011001
011010101101
101001001010
001101101010
100110111011
000000101111
111100000110
100000100000
011100010101
111000110010
111010001010
111000011010
000111001101
001111000000
011110000111
101001010110
011001001010
001010001001
111001001101
111111000001
001001111110
101000111110
001111010110
010000011111
010101010001
011001010010
111010000011
110000111100
110101000000
000111011110
101100101011
001110101011
000110100111
010100000001
101111000001
000110101100
111110010000
010000110101
000010000010
111111111101
011101001010
001111100101
100101010111
100011010100
110111110010
100110110011
011100101011
000001010011
010110101001
000111110101
101111001001
100010101110
111010010011
100101000111
111101110110
100101110011
001000110000
010111100001
100100011100
000100111011
101110111111
100110010001
100100111000
011010111010
111011101101
110010110100
101110101110
110001101101
110000011011
101000000010
010100000111
000110100101
011000101100
111001111110
101010100001
101111111001
100010010100
111001111000
100000011010
111000101110
010000000101
111111000111
100001011110
100110111001
101100101001
000011111101
111101110001
101010010011
101000110111
000000111111
110001000111
010110000010
011010101011
000100110000
101010111101
000100101101
001000001100
010000000111
100100100011
010010111101
000011001010
010100111110
011110110101
110001000110
010001100011
100111010100
100000101010
110101110001
111010000111
011000000011
001101011000
100110101110
101110101010
001101101001
111011101111
011010001110
111100001100
111111001000
111000110100
111101100011
010010101011
111101100101
111111000011
110100000011
110000010101
001111111000
100100101111
100101111000
011111110001
000110001000
001101110110
110010011010
110010101110
011010011010
000011000101
101101100000
100111111000
001100000010
101101101010
100100111001
111100101100
101100110010
010010111110
101001100101
100001111001
100010000101
110111010011
100000010011
111100000011
110000000010
101101001100
011001001000
101011110001
011111011010
110001010111
110001101100
011001000011
111101000011
110100100110
010111111001
011011111101
100100110011
011111000100
100010011110
000010011011
111101100111
111001111010
001111111100
111011001111
111011110100
000001101111
110010110110
101000101101
111000011100
001001101010
010110100001
000101110110
110101101101
000001001010
001011010011
010011001001
100101111010
000100011011
000100010010
101110110001
100000011001
011000010000
101111001010
101111100111
001000000110
100110011100
001010101011
111001001100
101110101011
100100011001
010011101001
110000001100
101010100000
011101010010
001010001110
001111101110
101101011100
011000100001
111100011101
010011110000
111010001011
011101000111
001100000100
101110010110
111101010111
011010100001
011110000101
011100001100
100100010011
000100110111
100101111110
010100011011
111110100000
011010011111
011100010011
010111101000
000000010111
100001110001
111100011100
001010110100
100111000000
001110101000
011010000001
110101110110
010010110101
011110100000
010101001001
011111111100
011100110010
100101011010
011011000101
101000110011
000000100101
110111101001
010101100110
101101111011
011100100110
101110011101
010000001001
000010100111
000000110110
101111101110
111100010011
001110101001
000000011000
110011100011
100001101010
010010101111
110111100011
011001110001
101000100001
011111001101
110101011101
000010110100
101110101111
111011000111
101101001001
100110000111
001110110100
001001100010
010110001001
100010001010
111100000101
101111101000
100001100110
111001010001
110101111100
000010100011
010001010010
011111101110
101110011010
110111101011
110101000110
101001111000
000110010010
110010011100
101101010100
000011001100
010100100000
001010001111
100111111110
100111011001
001011100010
101010101110
011011001111
010011111001
101101000110
100001111011
000100001000
111111010100
111001111101
001100101100
000010111000
000110110011
011101110011
111100000001
100000100101
100011001110
101000101100
101010010000
110110110000
010101011010
001101100101
010010101010
111000100100
110000001101
100001000000
010000100011
110001010001
010010111001
011100000110
100111100010
110110110110
100001000001
100010101010
001111100110
101101101110
110010010111
111100101011
101011110010
111001011110
010101000010
101011001010
001111110111
111100100101
111001110010
010110001101
110000110100
110010110000
101110101101
001010000001
101111000101
001110100001
010011101111
001100000111
101100110001
110011111011
011110001010
010111111101
000001111011
000001001001
001101000100
111110011011
010000011011
110011000110
001010000111
111100011110
000110111010
001011101110
001010000101
011101010111
110000101111
010010010011
110100110110
100011000011
001111100001
000110000000
000110111110
011001100001
010111111010
010101010111
000010011000
010011000010
100000011100
000111101101
001001010000
111000000011
010110001010
101000001000
010000001000
011001101110
110001110110
100010011011
011011001000
001100101111
011010111100
010100110011
000111000011
111101010011
000111011011
011110000110
101100000001
011111100000
111010101110
110001101010
100101000100
101011101001
110010101111
010010011011
100110111101
010000111010
010100011000
110111010101
001000011100
010011000110
101010001000
110010100011
001101011010
000011110000
100101001101
011001011111
100011100111
100001101000
001100011100
110001010000
101001110111
000001110010
010011100111
110011110011
101110111100
010110010101
101101101101
011001000010
001110010010
111111001101
010000100111
111101110011
011101101000
000111101100
110101011100
101001101101
101001100011
001101111011
000111111001
010000000010
100011110000
111111011110
011000110010
110001110010
010010100110
100000111101
010001001110
110001111111
110111000100
111010011100
001100101001
111001111100
001100101010
000000010010
010111101011
101100001111
110101001100
100101110010
011101111110
111100101110
010010000100
111010100010
001000101011
001000110111
100011001001
111101111010
110101010111
001101100010
000110011100
010111101111
111110000010
000100101011
010000100110
111000010100
111110011111
010001001100
000000001000
011001110101
101110100000
011100100001
110001100101
111101011001
101000101001
000011010100
001000001111
100101110100
001101110011
001000100111
101011011100
000100000010
001111011001
001001110101
001101011110
001111011010
001111001011
100010100011
000000010011
100100010100
111001000011
110110111110
001110011101
011011010111
011100001111
101110100101
111101010110
111100100111
101111101100
111000000001
101011011101
101011010101
100111000001
111100110001
111111100001
110001110101
011010111011
011100101000
111101011010
101010010100
111001011011
000100100010
111110000111
110110001000
011100111001
100100100100
010000100010
111010001000
010001011010
100010111110
001001001111
100001101101
000001111000
001100110110
000101101010
011000001001
111011101010
100110001101
110001010011
001000000111
110111001101
111010110111
100010101101
100101101011
001111100010
111011101001
000001110001
011110001000
001111100100
001001010010
110011100001
110011100100
111110111110
111011011001
111111111011
110011000111
011111000110
110001011101
010101111000
111100101101
010001101001
110001101001
111101000010
101110000111
000011101101
001110000010
011110011101
100011001101
100010111011
010001101011
011101110110
000010011001
100100010101
011100101101
010010100010
101110100011
100001111110
001011000000
000000000010
111000111110
100010010101
111101100010
100101000110
111001100100
101111001111
101110110111
010001000100
101101110010
111001101101
111101101111
000100000100
101011011011
111000110001
101010100101
110100000001
000111010110
001000010100
111110010111
011101000110
100111110010
010010000011
101011100100
101000100110
110101100001
001110000000
111101100100
011101010011
111100101000
100000010101
010011010011
011000001100
011111001111
110100001011
100000001111
001011001010
010001110110
101000110001
110001000101
010100111001
100000010010
101010111100
001110011001
110000100011
010010000111
101111110000
010101010010
011010001001
111010011111
000101100010
000101110011
100011111100
000100000111
101101100011
000101111010
000011010001
111000011101
100110000110
110100100011
010111000010
100011001011
010011100000
010111100101
100011100000
011000010010
010111010010
011010010110
111001111001
101100011100
110111111000
001001011000
000000011100
011000011010
011110000001
110110000100
010010100111
100000100111
011111110110
001010110000
110011000011
001010110011
011110101011
101011111001
010101011000
001100101011
000111100011
110010001111
001011101000
010101011111
001101000010
010011101000
111111110110
101011111111
110000001010
011011100110
110001111000
100001110000
100011000101
111001000101
101101101000
010000011110
100100111110
101010110010
000000101101
011100111011
101110101000
000011000001
100010110010
000101101101
110110100001
001001011111
101111110010
010011110101
101000110100
000001101000
111110000000
011000111100
010010010010
011110001111
110100111000
010011100011
000100100100
010000110001
000000001011
011110111000
010000000011
110010010100
111110111001
011111111001
010000101000
110101100011
100110001001
011011000111
100110101010
011011011111
011001011001
110110100101
011101000010
001000100100
000100110100
110000111010
000100100101
100000111110
010110110101
110100100001
110101100111
101000100000
101110001001
011000111110
110010101010
011100110000
001101000011
110101101110
111000001010
110101000001
100111101001
011010111001
010011110011
101000011011
111001010111
010000000001
100111001010
010101100111
011100101100
101011110111
100111100110
100111100000
110010111100
100011010011
110011110000
111100010010
001000101010
101110110110
010111101101
011110010000
000101000101
111110111000
111010100001
011101011101
111110000011
010011101010
010111000111
100010011001
000010000000
000000100100
110110110101
100010100100
101000000101
111111100000
101101101111
010100101110
111111100101
101000010101
001011111000
001110110011
010101100000
001010101010
111001100101
100110000011
010101011011
101001110100
100011100001
100110101000`
