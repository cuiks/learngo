package main

import "fmt"

func lengthOfNoneRepeatingSubStr(s string) int {
	maxLen := 0
	lastOccurred := make(map[rune]int)
	start := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfNoneRepeatingSubStr("abcadefc"))
}
