/*
Intersection of Two Arrays

Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.
*/

/*
INPUT: two arrays of integers
OUTPUT: a single array of integer

RULES
Explicit
- given two arrays of integers `nums1` and `nums2`,
  return an array of their intersection.
- Each element in the result must be unique and you
  return the result in any order.

Implicit
- Assume that a valid array will be provided
- Arrays of size 1 could be provided
- if arrays don't have intervals in common, return an empty array

T1
input: [1], [1]
output: [1]

T2
input: [1], [2]
output: []

T3
input: [1, 2, 2, 1], [2, 1]
output: [2]

1 2 2 1
  2 1

T4
input: [4, 9, 5], [9, 4, 9, 8, 4]
output: [9, 4], but [4, 9] is also accepted

Unique
  v            V
[ 4, 9, 5 ], [ 9, 4, 8, 1, 3 ]

[4, 9]



DATA STRUCTURE
- 2 maps with map[int]int
*/

package main

import (
	"fmt"
	"sort"
)

// T: O(N + M)
// S: O(1)
func intersection(nums1 []int, nums2 []int) []int {
	counts := make(map[int]bool)
	result := []int{}

	for _, num := range nums1 {
		counts[num] = true
	}

	for _, num := range nums2 {
		if counts[num] == true {
			counts[num] = false
			result = append(result, num)
		}
	}

	return result
}

// Time: O(N Log N + M Log M)
// Space: O(1)
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1) // Log N
	sort.Ints(nums2) // Log M

	i, j := 0, 0
	result := make([]int, 0)

	for i < len(nums1) && j < len(nums2) { // N
		num1, num2 := nums1[i], nums2[j]

		if num1 < num2 {
			i++
		} else if num1 > num2 {
			j++
		} else {
			// case where numbers are equal
			result = append(result, num1)

			for i < len(nums1) && nums1[i] == num1 {
				i++
			}

			for j < len(nums2) && nums2[j] == num1 {
				j++
			}
		}
	}

	return result
}

// A different way to sort arrays
// 	sort.Slice(nums1, func(i int, j int) bool {
// 		return nums1[i] < nums1[j]
// 	})

// 	sort.Slice(nums2, func(i int, j int) bool {
// 		return nums2[i] < nums2[j]
// 	})

func main() {
	fmt.Println(intersection([]int{1, 2, 3}, []int{4, 5, 6}))                                             // []
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))                                             // [2]
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))                                       // [9, 4], but [4, 9] is also accepted
	fmt.Println(intersection([]int{9, 3, 4, 3, 7, 3}, []int{4, 9, 4, 9, 4, 7, 8}))                        // [4, 7, 9]
	fmt.Println(intersection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // [4, 7, 9]
}
