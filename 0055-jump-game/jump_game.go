package main

import "fmt"

// Can we reach the last index of the array?
// each item in the array represents a maximum
// number of jumps we can take to move forward
// func canJump(nums []int) bool {

// 	// sub-problem
// 	n := len(nums)
// 	if n == 1 {
// 		return true
// 	}

// 	max := 0
// 	for index := 0; index < n-1 && max >= index; index++ {
// 		if max < index+nums[index] {
// 			max = index + nums[index]
// 		}

// 		if max >= n-1 {
// 			return true
// 		}
// 	}

// 	return false
// }

func canJump(nums []int) bool {

	for index, maxJumps := range nums {
		if index == len(nums)-1 {
			return true
		}

	}
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
}
