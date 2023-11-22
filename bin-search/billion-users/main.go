package main

import (
	"fmt"
	"math"
)

func main() {

	gr := []float64{1.5}
	fmt.Printf("GR: %+v, days: %d\n", gr, getBillionUsersDay(gr))

	gr = []float64{1.1, 1.2, 1.3}
	fmt.Printf("GR: %+v, days: %d\n", gr, getBillionUsersDay(gr))

	gr = []float64{1.01, 1.02}
	fmt.Printf("GR: %+v, days: %d\n", gr, getBillionUsersDay(gr))

}

func getBillionUsersDay(grs []float64) int {
	total := 0.0
	days := 0
	appus := make([]float64, len(grs))

	for ix := 0; ix < len(appus); ix++ {
		appus[ix] = 1
	}

	for day := 0; total < 1000000000; day++ {

		for ix := 0; ix < len(appus); ix++ {
			appus[ix] = math.Pow(grs[ix], float64(day))
		}

		total = 0
		for ix := 0; ix < len(appus); ix++ {
			total += appus[ix]
		}

		days = day
		if days > 100000 {
			break
		}
	}
	return days

}
