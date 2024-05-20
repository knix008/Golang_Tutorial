package main

import "fmt"

func getAverage(arr []int, size int) float32 {
	var i int
	var sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)
	return avg
}

func main() {
	var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("The average of {0, 1, 2, 3, 4, 5, 6, 7, 8, 9} is %f\n", getAverage(array, 10))
}
