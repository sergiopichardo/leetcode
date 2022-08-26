/*
162. Find Peak Element

- A peak element is an element that is strictly greater than its neighbors.
- Given a 0-indexed integer array nums, find a peak element, and return its index.
	If the array contains multiple peaks, return the index to any of the peaks.
- You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always
	considered to be strictly greater than a neighbor that is outside the array.
- You must write an algorithm that runs in O(log n) time.
*/

package main

import "fmt"

/*
Brute Force
Time: O(N)
Space: O(1)
*/
// func findPeakElement(nums []int) int {
// 	maxIndex := 0
// 	currentMax := 0

// 	for index, num := range nums {
// 		if num > currentMax {
// 			currentMax = num
// 			maxIndex = index
// 		}
// 	}
// 	return maxIndex
// }

/*
Optimized Solution
Time: O(Log N)
Space: O(1)
*/

/*
A peak occurs when a nunber is greater than
its left neighbor and its right neighbor
*/
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	// because we know we'll have a valid solution
	for left < right {
		// calculate mid index + protect for integer overflow
		mid := left + (right-left)/2
		// re-calculate boundaries
		// if the current element is smaller than its neighbor
		// [1 2 1 [3] 5 6 4]
		//  L 						R
		// [1 2 1 [3] 5 6 4]
		//          L 	R
		// [1 2 1 3 5 6 4]
		//
		//            LM R
		// [1 2 1 3 5 6 4]
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	result_1 := findPeakElement([]int{1, 2, 3, 1})          // index: 2
	result_2 := findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}) // index: 5

	fmt.Println(result_1)
	fmt.Println(result_2)
}
