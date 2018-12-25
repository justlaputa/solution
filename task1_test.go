package solution

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{0, 3, 3, 7, 5, 3, 11, 1}, // original test case 1
			7,                              // expect longest distance
		},
		{
			[]int{1, 4, 7, 3, 3, 5}, // original test case 2
			4,                       // expect longest distance
		},
		{
			[]int{0, 1, 2, 3, 4, 5}, // no duplicate, sequencial values
			1,
		},
	}

	for _, data := range testCases {
		result := Solution(data.input)
		if result != data.expected {
			t.Errorf("test failed on %v, expect: %d, got: %d", data.input, data.expected, result)
		}
	}
}
