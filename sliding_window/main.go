package main

import "fmt"

func main() {
	naiveApproach([]int{1, 2, 3, 4, 5}, 3)
	naiveApproach([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5)
}

// Given an array of integers, find the average of all contiguous subarrays of size ‘K’ in it.
// This is naive approach with O(N * K) time complexity
func naiveApproach(myArray []int, k int) {
	myAverageArray := make([]float32, 0)
	sum := 0

	for i := 0; i <= len(myArray)-k; i++ {
		sum = 0
		for j := 0; j < k; j++ {
			sum = sum + myArray[i+j]
		}
		result := float32(sum) / float32(k)
		myAverageArray = append(myAverageArray, result)
	}

	fmt.Println("---------------------------------------------")

	for i := 0; i < len(myAverageArray); i++ {
		fmt.Println(myAverageArray[i])
	}
}
