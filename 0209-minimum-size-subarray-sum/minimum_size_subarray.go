package main

/*
BRUTE FORCE
- find all subarrays
- then select the ones that have sum >= target
- return the minimal length

O(NlogN)

Binary search and then withing each iteration doing
something  O(N)

Or looping through the array and performing binary
search in each iteration

O(N) -> Pointers

anchor / runner (slow / fast pointers)
--------

1. Where does my anchor start? 0
2. Where does my runner start? 0
3. When do I have anchor? I have anchor when the sum reaches the target
4.
5.
6.

*/

func minSubArrayLen(target int, nums []int) int {

}

func main() {
	result_1 := minSubArrayLen()
	result_2 := minSubArrayLen()
	result_3 := minSubArrayLen()
}
