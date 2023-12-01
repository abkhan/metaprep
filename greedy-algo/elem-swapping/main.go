package main

import "fmt"


func findMinArray(arr []int, k int) []int {
 
  // find index of the smallest in the first k elements
  smallest := 10000001
  smix := -1
  for ix := 0; ix <= k; ix++ {
    if arr[ix] < smallest {
      smix = ix
      smallest = arr[ix]
    }
  }
  k = smix
  
  swaps(arr, k)
  return arr
}

func main() {

  arr := []int{5, 3, 1}
  k := 2
  fmt.Printf("for 5/3/1, k:%d >> resp: %+v\n", k, findMinArray(arr, k))
  
    arr = []int{8, 9, 11, 2, 1}
  k = 3
  fmt.Printf("for 8, 9, 11, 2, 1, k:%d >> resp: %+v\n", k, findMinArray(arr, k))
  
      arr = []int{8, 9, 11, 2, 1}
  k = 4
  fmt.Printf("for 8, 9, 11, 2, 1, k:%d >> resp: %+v\n", k, findMinArray(arr, k))
}

func swaps(a []int, k int) {
  fmt.Printf("swaps called with %d\n", k)
  if len(a) < k+1 {
    return
  }
  
  for ix := k; ix > 0; ix-- {
    a[ix], a[ix-1] = a[ix-1], a[ix]
  }
}

