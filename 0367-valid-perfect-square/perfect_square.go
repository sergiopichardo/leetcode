/*
367. Valid Perfect Square
Given a positive integer num, write a function which returns True if num is a perfect square else False.
Follow up: Do not use any built-in library function such as sqrt.
*/

/*
INPUT: number
OUTPUT: boolean

RULES
- Given a positive integer number, write a function which return True
  if the number is a perfect square. Otherwise, return False.
- DO NOT use any built-in library functions such as `sqrt`.

Questions?
- What is a perfect square?
- How to calculate the perfect square of a number?

*/

package main

import "fmt"

/*
Time complexity: O(N)
Space complexity: O(1)
*/
// func isPerfectSquare(num int) bool {
// 	if num == 1 {
// 		return true
// 	}

// 	for factor := 1; factor < num; factor++ {
// 		square := factor * factor
// 		if square == num {
// 			return true
// 		}
// 	}

//		return false
//	}

/*
Time: O(Log N)
Space: O(1)
*/
func isPerfectSquare(num int) bool {
	high := num
	low := 1

	for low <= high {
		mid := low + (high-low)/2
		square := mid * mid

		if square == num {
			return true
		} else if square < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func main() {
	result_1 := isPerfectSquare(25)
	result_2 := isPerfectSquare(100)
	result_3 := isPerfectSquare(36)
	result_4 := isPerfectSquare(49)
	result_5 := isPerfectSquare(29)
	result_6 := isPerfectSquare(1000)

	fmt.Println(result_1)
	fmt.Println(result_2)
	fmt.Println(result_3)
	fmt.Println(result_4)
	fmt.Println(result_5)
	fmt.Println(result_6)
}
