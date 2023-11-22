// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "strings"
import "fmt"

func minLengthSubstring(s string, t string) int {

	minX := len(s) + 1
	maxX := -1

	for _, c := range t {
		if ci := strings.Index(s, string(c)); ci == -1 {
			return -1
		} else {
			if ci > maxX {
				maxX = ci
			}
			if ci < minX {
				minX = ci
			}
		}
	}
	return (maxX - minX) + 1
}

func main() {
	strs := "dcbefebce"
	strt := "fd"
	fmt.Printf("For %s, subs %s, ret: %d\n", strs, strt, minLengthSubstring(strs, strt))

	strs = "dcbefebce"
	strt = "zy"
	fmt.Printf("For %s, subs %s, ret: %d\n", strs, strt, minLengthSubstring(strs, strt))

	strs = "abcde"
	strt = "edcba"
	fmt.Printf("For %s, subs %s, ret: %d\n", strs, strt, minLengthSubstring(strs, strt))

	strs = "abcdefghijklmonpqurstuvwzmynothingbutwhatcreatorsayz"
	strt = "azy"
	fmt.Printf("For %s, subs %s, ret: %d\n", strs, strt, minLengthSubstring(strs, strt))
}
