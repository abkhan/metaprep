// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func areTheyEqual(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		fmt.Println("arrays not equal")
		return false
	}
	if len(a1) == 0 && len(a2) == 0 {
		fmt.Println("arrays nil")
		return true
	}
	if len(a1) == 1 && len(a2) == 1 && a1[0] == a2[0] {
		fmt.Println("array len 1")
		return true
	}

	if verifyEq(a1, a2) {
		fmt.Println("array same")
		return true
	}

	diffmap := diffs(a1, a2)
	for k, v := range diffmap {
		fmt.Printf("> Start:%d, Len:%d for %+v\n", k, v, a2)
		reverseSub(a2, k, v)
	}

	if verifyEq(a1, a2) {
		return true
	}

	return false
}

func main() {

	a1 := []int{1, 2, 3, 4}
	a2 := []int{1, 4, 3, 2}
	fmt.Printf("Eq: %+v for %+v, %+v\n\n", areTheyEqual(a1, a2), a1, a2)

	a1 = []int{3, 50, 60, 70, 90}
	a2 = []int{50, 3, 60, 70, 90}
	fmt.Printf("Eq: %+v for %+v, %+v\n\n", areTheyEqual(a1, a2), a1, a2)

	a1 = []int{111, 122, 155, 199, 333, 4444, 5555, 1}
	a2 = []int{111, 199, 155, 122, 333, 5555, 4444, 1}
	fmt.Printf("Eq: %+v for %+v, %+v\n\n", areTheyEqual(a1, a2), a1, a2)

	a1 = []int{111, 122, 156, 199, 333, 4444, 5555, 1}
	a2 = []int{111, 199, 155, 122, 333, 5555, 4444, 1}
	fmt.Printf("Eq: %+v for %+v, %+v\n\n", areTheyEqual(a1, a2), a1, a2)
}

// verifyEq, verify the contents are equal,
// the lenfth of a1 and a2 must be equal
func verifyEq(a1 []int, a2 []int) bool {
	for ix := 0; ix < len(a1); ix++ {
		if a1[ix] != a2[ix] {
			return false
		}
	}
	return true
}

// reverseSub would reverse a sub-array from within array
// does not check index validity
func reverseSub(a []int, start, alen int) {
	//fmt.Printf("revSub st:%d, len:%d for %+v\n", start, alen, a)
	from := start
	to := start + alen - 1
	for ; from < to; from, to = from+1, to-1 {
		temp := a[from]
		a[from] = a[to]
		a[to] = temp
	}
}

func diffs(a1 []int, a2 []int) map[int]int {
	rmap := map[int]int{}

	pdiff := false
	startv := 0
	startx := 0
	alen := 0
	for ix := 0; ix < len(a1); ix++ {

		if pdiff { // match startv
			alen++
			if a2[ix] == startv {
				pdiff = false
				rmap[startx] = alen
				startx = 0
				alen = 0
			}

			continue
		}

		if a1[ix] != a2[ix] {
			if !pdiff {
				pdiff = true
				startv = a1[ix]
				startx = ix
				alen = 1
			}
		}
	}

	return rmap
}
