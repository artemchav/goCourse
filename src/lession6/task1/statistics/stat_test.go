package statistic

import "testing"

func TestSum(t *testing.T) {
	type Case struct {
		input []int
		res   int
	}
	cases := []Case{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{5, 2, 3, 4, 5}, 19},
		{[]int{1, 3, 3, 4, 5}, 16},
		{[]int{1, 10, 3, 4, 5}, 23},
	}

	for _, v := range cases {
		sum := Sum(v.input...)
		if sum != v.res {
			t.Errorf("Expected %d, got %d", v.res, sum)		}
	}
}
