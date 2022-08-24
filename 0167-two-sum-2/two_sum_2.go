package main

import "fmt"

func twoSum(nums []int, target int) []int {
	leftPointer := 0
	rightPointer := len(nums) - 1
	result := []int{}

	for {
		sum := nums[leftPointer] + nums[rightPointer]

		if sum == target {
			result = append(result, leftPointer+1, rightPointer+1)
			break
		} else if sum > target {
			rightPointer--
		} else {
			leftPointer++
		}
	}

	return result
}

func main() {
	result_1 := twoSum([]int{2, 7, 11, 15}, 9)
	result_2 := twoSum([]int{2, 3, 4}, 6)
	result_3 := twoSum([]int{-1, 0}, -1)

	fmt.Println(result_1)
	fmt.Println(result_2)
	fmt.Println(result_3)
}
