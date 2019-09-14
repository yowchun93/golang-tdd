package sum

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, n := range numbersToSum {
		sums = append(sums, Sum(n))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, n := range numbersToSum {
		sums = append(sums, n[len(n)-1])
	}
	return sums
}
