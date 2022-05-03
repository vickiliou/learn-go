package arrays

// Sum takes a slice of integers and returns the sum of them.
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// SumAll takes a varying number of slices and returns the sums of every slice passed in.
func SumAll(numToSum ...[]int) []int {
	var sums []int

	for _, num := range numToSum {
		sums = append(sums, Sum(num))
	}
	return sums
}
