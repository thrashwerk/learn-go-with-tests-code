package main

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }

	return Reduce(numbers, add, 0)
}

func SumAll(numberLists ...[]int) []int {
	var sums []int

	for _, l := range numberLists {
		sums = append(sums, Sum(l))
	}

	return sums
}

func SumAllTails(numberLists ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numberLists, sumTail, []int{})
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue

	for _, x := range collection {
		result = f(result, x)
	}

	return result
}

func Find[A any](collection []A, f func(A) bool) (A, bool) {
	var zero A
	for _, v := range collection {
		if f(v) {
			return v, true
		}
	}

	return zero, false
}
