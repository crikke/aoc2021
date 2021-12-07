package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	fuel, _ := day7(testinput)

	assert.Equal(t, 37, fuel)
}

// mean fungerar += 1, gör bättre lösning när jag har tid
func TestPart2(t *testing.T) {
	_, fuel := day7(testinput)

	assert.Equal(t, 168, fuel)
}

var testinput = "16,1,2,0,4,2,7,1,2,14"
