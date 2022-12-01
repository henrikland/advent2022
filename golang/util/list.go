package util

func ListSum(list []int) int {
	sum := 0

	for _, v := range list {
		sum += v
	}

	return sum
}
