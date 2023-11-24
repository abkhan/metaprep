package main

import "fmt"

func canGetExactChange(targetMoney int, denominations []int) bool {
 
  for ix := len(denominations) -1; ix >= 0; ix-- {
    targetMoney = targetMoney % denominations[ix]
  }
  if targetMoney == 0 {
    return true
  }
  return false
}

func main() {
  d := []int{5, 10, 25, 100, 200}
  t := 94
  fmt.Printf("For Target [%d], and den[%+v], exactChange possible? %+v\n", t, d, canGetExactChange(t, d))
  
  d = []int{4, 17, 29}
  t = 75
  fmt.Printf("For Target [%d], and den[%+v], exactChange possible? %+v\n", t, d, canGetExactChange(t, d))

  d = []int{5, 25, 50, 100, 500, 1000, 2000, 5000, 10000}
  t = 2935
  fmt.Printf("For Target [%d], and den[%+v], exactChange possible? %+v\n", t, d, canGetExactChange(t, d))

  d = []int{5, 25, 50, 100, 500, 1000, 2000, 5000, 10000}
  t = 750545
  fmt.Printf("For Target [%d], and den[%+v], exactChange possible? %+v\n", t, d, canGetExactChange(t, d))

}
