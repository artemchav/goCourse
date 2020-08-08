package statistic

func Sum(elems ...int) int {
	sum := 0
	for _, v := range elems {
		sum += v
	}

	return sum
}
