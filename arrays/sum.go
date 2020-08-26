package main

// Sum takes an array of integers and returns their sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll calculates the respective sums of every slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	numberOfGroups := len(numbersToSum)
	sums := make([]int, numberOfGroups)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
