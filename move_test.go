package gomovearrposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7}

	testCases := []struct {
		k      int
		output []int
	}{
		{1, []int{1, 2, 3, 4, 5, 6, 7}},
		{2, []int{1, 2, 3, 4, 5, 7, 6}},
		{3, []int{1, 2, 3, 4, 7, 6, 5}},
		{4, []int{1, 2, 3, 7, 6, 5, 4}},
		{5, []int{1, 2, 7, 6, 5, 4, 3}},
		{6, []int{1, 7, 6, 5, 4, 3, 2}},
		{7, []int{7, 6, 5, 4, 3, 2, 1}},
		{100, []int{7, 6, 5, 4, 3, 2, 1}},
		{-100, []int{1, 2, 3, 4, 5, 6, 7}},
		{8732438432743, []int{7, 6, 5, 4, 3, 2, 1}},
		{-483243720987532, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, v := range testCases {
		assert.Equal(t, MoveByCountFront(input, v.k), v.output, "I hope this passed bruh:(")
	}

}
