package main

import "fmt"

func getMilestoneDays(revenues []int, milestones []int) []int {
	resp := make([]int, len(milestones))

	for ix := 0; ix < len(milestones); ix++ {
		resp[ix] = oneMile(revenues, milestones[ix])
	}

	return resp
}

func main() {
	r := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	miles := []int{100, 200, 500}
	fmt.Printf("Rev:%+v, Miles:%+v >> OP:%+v\n", r, miles, getMilestoneDays(r, miles))

	r = []int{40, 44, 48, 50, 60, 66, 70, 78, 80, 88, 90, 100}
	miles = []int{200, 250, 500, 1000}
	fmt.Printf("Rev:%+v, Miles:%+v >> OP:%+v\n", r, miles, getMilestoneDays(r, miles))

	r = []int{40, 44, 48, 50, 60, 66, 70, 78, 86, 88, 90, 100}
	miles = []int{400, 500, 600, 900}
	fmt.Printf("Rev:%+v, Miles:%+v >> OP:%+v\n", r, miles, getMilestoneDays(r, miles))
}

func oneMile(revenues []int, milestone int) int {
	cumu := 0
	for ix := 0; ix < len(revenues); ix++ {
		cumu += revenues[ix]
		if cumu >= milestone {
			return ix + 1
		}
	}
	return -1
}
