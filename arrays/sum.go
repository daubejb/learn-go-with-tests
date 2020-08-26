package sum

// Sum takes an array of integers and returns their sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
