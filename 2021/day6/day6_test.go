package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	sum := part1(18, testinput)

	assert.Equal(t, 26, sum)
}

var testinput = "3,4,3,1,2"
