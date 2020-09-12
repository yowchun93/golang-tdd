package sum

// Sum adds all the numbers together in the array
func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	// return sum

	// using range instead of for loop
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, number := range numbersToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
