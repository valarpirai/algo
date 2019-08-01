package main

import (
	"reflect"
	"testing"
)

func TestGetProductOfOthers(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 7, 3, 4}, []int{84, 12, 28, 21}},
		{[]int{2, 4, 10}, []int{40, 20, 8}},
		{[]int{2, 4, 0}, []int{0, 0, 8}},
	}

	for _, tt := range tests {
		result := getProductOfOthers(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

// getProductOfOthers takes a list of numbers and returns a corresponding list
// where every index holds the product of every other values except the value
// in that index. We make two passes through our input list so the time
// complexity is O(n). We return a list of products as a result so the space
// complexity is O(n).
func getProductOfOthers(list []int) []int {
	if len(list) < 2 {
		return []int{}
	}

	out := make([]int, len(list))

	// get product of all other numbers before their indexes.
	start1 := 1
	for i := 0; i < len(list); i++ {
		out[i] = start1
		start1 *= list[i]
	}

	// get product of all other numbers after their indexes then multiply them
	// with their corresponding values.
	start2 := 1
	for i := len(list) - 1; i > -1; i-- {
		out[i] *= start2
		start2 *= list[i]
	}

	return out
}