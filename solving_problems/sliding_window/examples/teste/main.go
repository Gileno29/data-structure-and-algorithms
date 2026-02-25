package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	var s1Count, s2Count [26]int
	for i := 0; i < n1; i++ {
		s1Count[s1[i]-'a']++ // 'a' is the ascii 97
	}

	// Initial window of size n1
	for i := 0; i < n1; i++ {
		s2Count[s2[i]-'a']++
	}

	// Slide the window
	for i := 0; i < n2-n1; i++ {
		fmt.Println(s1Count, s2Count)
		// 1. Check if current window matches
		if s1Count == s2Count {
			return true
		}

		// 2. Prepare for NEXT window:
		// Add the character at (i + n1)
		s2Count[s2[i+n1]-'a']++
		// Remove the character at (i)
		s2Count[s2[i]-'a']--
	}

	// Check the very last window
	return s1Count == s2Count
}

func main() {
	str1 := "abc"
	str2 := "lecabee"

	fmt.Print(checkInclusion(str1, str2))
}
