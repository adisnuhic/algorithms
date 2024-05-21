package main

import (
	"fmt"
)

func main() {
	naiveApproach([]int{1, 2, 3, 4, 5}, 3)              // [2.0, 3.0, 4.0]
	naiveApproach([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5) // [2.2, 2.8, 2.4, 3.6, 2.8]

	sliding_window_approach([]int{1, 2, 3, 4, 5}, 3)              // [2.0, 3.0, 4.0]
	sliding_window_approach([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5) // [2.2, 2.8, 2.4, 3.6, 2.8]
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

// This is sliding window approach with O(N) time complexity
func sliding_window_approach(myArray []int, k int) {
	myAverageArray := make([]float32, 0)
	windowStart := 0
	windowSum := 0

	for windowEnd := 0; windowEnd < len(myArray); windowEnd++ {
		windowSum += myArray[windowEnd]

		if windowEnd >= k-1 {
			// calculate the average
			myAverageArray = append(myAverageArray, float32(windowSum)/float32(k))

			// remove the element going out of the window
			windowSum -= myArray[windowStart]

			// slide the window ahead
			windowStart++
		}
	}

	fmt.Println("------------------- Sliding Window --------------------------")

	for i := 0; i < len(myAverageArray); i++ {
		fmt.Println(myAverageArray[i])
	}

}
