package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	fuel := part1(testinput)

	assert.Equal(t, 37, fuel)
}

var testinput = "16,1,2,0,4,2,7,1,2,14"
