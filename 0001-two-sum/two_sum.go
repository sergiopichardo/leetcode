package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	result := []int{}

	for index, value := range nums {
		complement := target - value

		if _, ok := cache[complement]; ok {
			result = append(result, cache[complement], index)
		} else {
			cache[value] = index
		}
	}

	return result
}

func main() {
	result1 := twoSum([]int{2, 11, 15, 7, 38}, 9) // [0, 3]
	result2 := twoSum([]int{39, 9, 22, 2, -3}, 6) // [1, 4]
	fmt.Println(result1)
	fmt.Println(result2)
}
