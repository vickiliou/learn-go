package arrays

// Sum takes a slice of integers and returns the sum of them.
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// SumAllTails takes a varying number of slices and returns the sums of every slice passed in, except the first number given.
func SumAllTails(numToSum ...[]int) []int {
	var sums []int

	for _, num := range numToSum {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
