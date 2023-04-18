package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	symvols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int = 0
	var sym byte = ' '
	for i := len(s)-1; i >= 0; i-- {
		if symvols[s[i]] < result && sym != s[i] {
			result = result - symvols[s[i]]
			sym = s[i]
		} else {
			result = result + symvols[s[i]]
			sym = s[i]
		}
	}
	return result
}