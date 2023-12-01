package main

import "fmt"
import "sort"

func findMedian(arr []int) []int {

  resp := []int{}
  if len(arr) < 2 {
    return arr
  }
  resp = append(resp, arr[0])
  for ix := 1; ix < len(arr); ix++ {
    resp = append(resp, findMed(arr[:ix+1]))
  }
  
  return resp
}

func main() {
  arr := []int{5, 15, 1, 3}
  fmt.Printf("Input: %+v >>> Resp: %+v\n", arr, findMedian(arr))

    arr = []int{1, 2}
  fmt.Printf("Input: %+v >>> Resp: %+v\n", arr, findMedian(arr))

    arr = []int{9, 8, 6, 7, 5, 4, 2, 1, 1, 2, 3}
  fmt.Printf("Input: %+v >>> Resp: %+v\n", arr, findMedian(arr))
}

func findMed(a []int) int {
  if len(a) < 1 {
    return -1
  }
  sort.Ints(a)
  if len(a) %2 == 1 {
    fmt.Printf("  >> [%+v], ret: %d\n", a, a[(len(a)/2)])
    return a[(len(a)/2)]
  }
  retval := ((a[(len(a)/2) -1] + a[len(a)/2])) / 2
  fmt.Printf("  >> [%+v], ret: %d\n", a, retval)
  
  return retval
}
~
