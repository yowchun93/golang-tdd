package sum

// Sum returns the sum of the array
// func Sum(num [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += num[i]
// 	}

// 	return sum
// }

// Sum uses range for iteration
func Sum(num []int) int {
	sum := 0
	for _, number := range num {
		sum += number
	}

	return sum
}

// SumAll sums each arrays
func SumAll(numbersToSum ...[]int) []int {
	a := []int{}
	for _, arr := range numbersToSum {
		a = append(a, Sum(arr))
	}
	return a
}

func SumAllTails(numbersToSum ...[]int) []int {
	a := []int{}
	for _, arr := range numbersToSum {
		a = append(a, Sum(arr[1:]))
	}
	return a
}
