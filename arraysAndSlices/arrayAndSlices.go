package arraysandslices

func ArraySum(numbers [5]int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SlicesSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SlicesSum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	if len(numbersToSum) == 0 {
		return append(sums, 0)
	} else {
		for _, numbers := range numbersToSum {
			if len(numbers) == 0 {
				sums = append(sums, 0)
			} else {
				sums = append(sums, SlicesSum(numbers[1:]))
			}
		}
	}
	return sums
}