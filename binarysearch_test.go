package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 3, -1},
		{[]int{2}, 2, 0},
		{[]int{2}, 3, -1},

		// Add more test cases as needed
	}

	for _, c := range cases {
		got := binarySearch(c.arr, c.target)
		if got != c.want {
			t.Errorf("binarySearch(%v, %d) == %d, want %d", c.arr, c.target, got, c.want)
		}
	}
}
