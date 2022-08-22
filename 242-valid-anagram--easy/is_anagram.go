package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[string]int)
	balance := 0

	for _, letter := range s {
		if _, ok := counter[string(letter)]; ok {
			balance += 1
		}
	}

	for _, compareletter := range t {
		// if _, ok := counter[string(compareletter)]; ok {
		// 	counter[string(compareletter)] -= 1
		// }
		counter[string(compareletter)] -= 1
	}

	for key, _ := range counter {
		if counter[string(key)] != 0 {
			return false
		}
	}

	return true
}

func main() {
	result1 := isAnagram("anagram", "nagaram")
	result2 := isAnagram("anagram", "nagarams")
	result3 := isAnagram("anagram", "wadddup")
	result4 := isAnagram("aacc", "ccac")
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
