package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	fuel, _ := day7(testinput)

	assert.Equal(t, 37, fuel)
}

// dagens var inte jättebra..
// testdatan ska rundas uppåt mean = 4.9 ska bli 5
// medans riktiga ska rundas nedåt, 464.567 blir 464
func TestPart2(t *testing.T) {
	_, fuel := day7(input)

	assert.Equal(t, 168, fuel)
}

var testinput = "16,1,2,0,4,2,7,1,2,14"
