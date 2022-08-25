package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	runner := 0
	left := 1
	right := len(nums) - 1
	result := [][]int{}

	sort.Ints(nums)

	if len(nums) == 0 {
		return result
	}

	for runner != len(nums)-1 {
		fmt.Println("triplet:", nums[runner], nums[left], nums[right])

		sum := nums[runner] + nums[left] + nums[right]
		fmt.Println("sum:", sum)

		if (runner != left && runner != right && left != right) && sum == 0 {
			result = append(result, []int{nums[runner], nums[left], nums[right]})
		}

		if sum > 0 {
			left++
		} else if sum < 0 {
			right--
		}

		runner++
	}

	return result
}

func main() {
	nums_1 := []int{-3, 3, -4, -3, 0, 2} // [[-1,-1,2],[-1,0,1]]
	// nums_2 := []int{-1, 0, 1, 2, -1, -4} // [[-1, 0, 1], [0, 1, -1]]
	// nums_3 := []int{0, 1, 1}             // [[]]
	// nums_4 := []int{0, 0, 0}             // [[0,0,0]]

	result_1 := threeSum(nums_1)
	// result_2 := threeSum(nums_2)
	// result_3 := threeSum(nums_3)
	// result_4 := threeSum(nums_4)

	fmt.Println(result_1)
	// fmt.Println(result_2)
	// fmt.Println(result_3)
	// fmt.Println(result_4)
}

/*
Line 31: Char 9: cannot use ret (type []int) as type [][]int in return argument (solution.go)
*/
