package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	substringMap := make(map[uint8]int)
    var length int = len(s)
	var result int = 0
	var i int = 0
	for j:= 0; j < length; j++ {
		if val, ok := substringMap[s[j]]; !ok {
			substringMap[s[j]] = j
		} else {
			if j - i > result {
				result = j - i
			}
			for k := i; k <= val; k++ {
				delete(substringMap, s[k])
			}
			i = val + 1
			substringMap[s[j]] = j
		}
	}
	if length - i > result {
		result = length - i
	}
	return result
}