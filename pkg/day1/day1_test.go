package day1_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetLargerMeasurementsCount(t *testing.T) {
	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/1/input")

	assert.NoError(t, err)

}
