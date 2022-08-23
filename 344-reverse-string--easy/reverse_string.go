package main

import "fmt"

/*
INPUT: []byte
OUTPUT: []byte

RULES
Explicit
- Write a function that reverses a string
- The input string is given a slice of characters `s`
- You must do this by modifying the input array in-place with O(1) extra memory
- s[i] is a printable ascii character
- s.length >= 1
- s.length <= 10^5

Implicit
- Assume we'll always get a valid string `s` of at least size `1`
- It's ok to mutate the input slice

EXAMPLES
T1: odd length
input: ["h", "e", "l", "l", "o"]
output: ["o", "l", "l", "h"]

T2: even length
input: ["s", "t", "a", "r", "l", "o", "r", "d"]
output: ["d", "r", "o", "l", "r", "a", "t", "s"]

DATA STRUCTURE
- NA we'll be reversing the string in place
	We'll use the given array

Questions
- How far do we need to iterate?
	- math.Floor(s.length / 2)

- The 2 indices we're swapping the letters with
	- left index starts at 0 and is incremented by 1
	- right index starts at s.length-1 and we decrement by 1


ALGORITHM
- given the slice of runes `s`

- set leftPointer  is 0
- set rightPointer is math.Floor(s.length / 2)

- while leftPointer is not equal to math.Floor(s.length / 2)
	- swap the characters at leftPointer with character at right pointer
	- increment left pointer by 1
	- decrement left pointer by 1

- return s

*/


// Time complexity: O(N)
// Space complexity: O(1) 
func reverseString(s []byte) []byte {
	leftPointer := 0
	rightPointer := len(s) - 1

	for leftPointer != len(s)/2 {
		temp := s[leftPointer]
		s[leftPointer] = s[rightPointer]
		s[rightPointer] = temp

		rightPointer--
		leftPointer++
	}

	return s
}

func main() {
	result1 := reverseString([]byte("starlord"))
	result2 := reverseString([]byte("hello"))

	fmt.Println(string(result1))
	fmt.Println(string(result2))
}

