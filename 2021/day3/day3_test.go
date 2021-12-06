package main

import (
	"fmt"
	"strconv"
	"strings"
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

	gamma, epsilon := part1(testinput)

	assert.Equal(t, 22, gamma)
	assert.Equal(t, 9, epsilon)
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

// fan va coolt, har lyckats implementera en sorterinsalgoritm :-D
// problemet just nu är att den root är den största biten.
// Detta gör att siffor med olika antal bits hamnar fel.
// Exempel:
// 101  8
// 10   4
// för att det ska gå så behöver 10 ha skrivas om till 010
func TestSort(t *testing.T) {
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

	input := strings.Split(testinput, "\n")

	root := node{}
	for _, item := range input {
		i, err := strconv.ParseInt(item, 2, 0)
		bitCount := len(item)

		if err != nil {
			panic(err)
		}
		fmt.Printf(" %d", i)
		root.insert(int(i), bitCount)
	}

	fmt.Println()
	res := root.Out()

	for i := 1; i < len(res); i++ {
		assert.GreaterOrEqual(t, res[i], res[i-1])
	}

	fmt.Print(res)
	fmt.Println()
}
