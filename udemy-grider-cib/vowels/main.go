package main

import "strings"

func main() {

}

func vowels(str string) int {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0
	str = strings.ToLower(str)

	for _, s := range str {
		if vowels[s] {
			count++
		}
	}
	return count

}
