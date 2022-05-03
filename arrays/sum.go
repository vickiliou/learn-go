package arrays

// Sum takes a slice of integers and returns the sum of them.
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
