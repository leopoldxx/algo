package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubarray2 xxx
func TestSubarray2(t *testing.T) {
	testCases := []struct {
		nums   []int
		min    int
		max    int
		expect int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			min:    1,
			max:    3,
			expect: 4,
		},
	}

	for _, test := range testCases {
		result := SubArray2(test.nums, test.min, test.max)
		log.Println(result)
		assert.Equal(t, test.expect, result, "should equal")
	}
}
