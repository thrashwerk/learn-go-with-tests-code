package main

func Sum(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(numberLists ...[]int) []int {
	var sums []int

	for _, l := range numberLists {
		sums = append(sums, Sum(l))
	}

	return sums
}

func SumAllTails(numberLists ...[]int) []int {
	var sums []int

	for _, l := range numberLists {
		if len(l) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := l[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
