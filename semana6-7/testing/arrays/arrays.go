package arrays

func Sum(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return res
}

func SumAll(slices ...[]int) []int {
	res := make([]int, len(slices))

	for i, slice := range slices {
		res[i] = Sum(slice)
	}

	return res
}

func SumAllTails(slices ...[]int) []int {
	res := make([]int, len(slices))

	for i, slice := range slices {
		if len(slice) == 0 {
			continue
		}
		res[i] = Sum(slice[1:])
	}

	return res
}
