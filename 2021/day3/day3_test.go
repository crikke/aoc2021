package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	testinput := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	gamma, epsilon := day3(testinput)

	assert.Equal(t, int64(22), gamma)
	assert.Equal(t, int64(9), epsilon)
}

func TestPart2(t *testing.T) {
	testinput := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	oxygen, co2 := part2(testinput)

	assert.Equal(t, oxygen, 23)
	assert.Equal(t, co2, 10)
}
