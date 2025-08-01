package main

func SumWithForLoop(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func SumWithWhileStyle(n int) int {
	sum := 0
	i := 1
	for i <= n {
		sum += i
		i++
	}
	return sum
}

func SumSliceValues(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
