/*
https://leetcode.com/problems/merge-sorted-array/


You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func merge(nums1 []int, m int, nums2 []int, n int) {

}

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*/

/*
INPUT: []int, []int, int, int
OUTPUT: []int

RULES
Explicit
- given two arrays of integers in ascending order (already sorted)
- merge nums1 and num2 into a single array
- mutate nums1
- m: is the non-zero numbers in nums1 that should be merged (the numbers could be also be zero)
- n: last elements set to 0 and should be ignored
	- nums2 has a size of n

EXAMPLES
T1      f              s
input: [1,2,3,0,0,0], [2,5,6], m: 3, n: 3
output: [1,2,2,3,5,6]

T2
input: [3,4,7,0,0,0], [2,5,6], m: 3, n: 3
output: [2, 3, 4, 5, 6, 7]

T3


[2, 3, 5, 5, 6]

            f                       s
input: [3,5,11, 12, 13, 14,0,0,0], [2,5,6], m: 3, n: 3

output: [2, 3, 5, 5, 6, 11]

for i, v := range result {
	nums1[i] = v
}

Cases
-[0]  [2] m=0  n=1
- [] [] m=0 n=0

DATA STRUCTURE
- slice

ALGORITHM
- define first and set 0
- define second pointer and set to 0
- define temporary and set to empty slice

- while `first` < len(nums1) - m AND `second` < len of nums2
	- compare the value at `first` with value at `second`
		- if value at `first` < value at `second`
			- push to the temporary array
			- increment `second` by 1
		- otherwise, if value at `first` is > value at `second`
			- push to the temporary array
			- increment `first` by 1

*/

package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

}

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}

	merge(arr1, 3, arr2, 3)
	fmt.Println(arr1) //[1, 2, 2, 3, 5, 6]
}
