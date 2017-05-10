package slice

import "fmt"

const testVersion = 1

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var stringSet []string
	for i := 0; i <= len(s)-n; i++ {
		stringSet = append(stringSet, s[i:i+n])
		fmt.Println(stringSet)
	}
	return stringSet
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

// First returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	fmt.Println(n > len(s))
	if n > len(s) {
		return
	}
	return s[:n], true
}
