package main

import "fmt"

func numberOfWays(arr []int, k int) int {

  all := pairs(arr, k)
  return len(all)
}

func main() {
  k := 6
  arr := []int{1, 2, 3, 4, 3}
  fmt.Printf("Arr: %+v, Sum:%d, total:%d\n", arr, k, numberOfWays(arr, k))

    k = 6
  arr = []int{1, 5, 3, 3, 3}
  fmt.Printf("Arr: %+v, Sum:%d, total:%d\n", arr, k, numberOfWays(arr, k))

    k = 7
  arr = []int{1, 2, 4, 3, 3, 3, 5, 5, 5}
  fmt.Printf("Arr: %+v, Sum:%d, total:%d\n", arr, k, numberOfWays(arr, k))

}


func pairs(arr []int, k int) [][2]int {
  resp := [][2]int{}
  for ix := 0; ix < len(arr)-1; ix++ {
    for jx := ix + 1; jx < len(arr); jx++ {
      if arr[ix] + arr[jx] == k {
        resp = append(resp, [2]int{ix, jx})
      }
    }
  }
  return resp
}

