package main

import "fmt"

func findSignatureCounts(arr []int) []int {
	sign := make([]int, len(arr))
	current := make([]int, len(arr))

	for ix := 0; ix < len(arr); ix++ {
		ybsign(ix, arr, current, sign)
		current = pass(arr, current)
		//fmt.Printf("Pass:%d >> Current %+v\n", ix, current)
	}

	return sign
}

func main() {
	testa := []int{1, 2}
	fmt.Printf("Ret %+v for input %+v\n\n", findSignatureCounts(testa), testa)

	testa = []int{2, 1}
	fmt.Printf("Ret %+v for input %+v\n\n", findSignatureCounts(testa), testa)

	testa = []int{4, 2, 3, 1}
	fmt.Printf("Ret %+v for input %+v\n\n", findSignatureCounts(testa), testa)

	testa = []int{5, 1, 3, 4, 2}
	fmt.Printf("Ret %+v for input %+v\n\n", findSignatureCounts(testa), testa)

	testa = []int{5, 6, 7, 1, 2, 3, 4}
	fmt.Printf("Ret %+v for input %+v\n\n", findSignatureCounts(testa), testa)
}

// sign calculates the signature on each yearbook
func ybsign(pass int, passing []int, current []int, sign []int) {
	if pass == 0 {
		for ix := 0; ix < len(sign); ix++ {
			sign[ix]++
			current[ix] = ix
		}
		return
	}

	for ix := 0; ix < len(sign); ix++ { // sign the book at the index
		if current[ix] != ix { // only if not the same as index
			student := current[ix]
			sign[student]++
		}
	}
}

// pass uses the pass-array to move the yearbooks
func pass(passmap []int, current []int) []int {
	newcur := make([]int, len(current))
	for ix := 0; ix < len(passmap); ix++ {
		// ix is the student, value of the current is the student id
		next := passmap[ix]        // where to sent this to
		next -= 1                  // 0 based index
		newcur[next] = current[ix] // take from current and move to new current
	}
	return newcur
}
