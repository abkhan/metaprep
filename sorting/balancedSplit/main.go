package main

import "fmt"
import "sort"

func balancedSplitExists(arr []int) bool {
  sort.Ints(arr)

  totalval := 0
  for ix :=0; ix < len(arr); ix++ {
    totalval += arr[ix]
  }
  if totalval % 2 != 0 {
    fmt.Printf("total [%d] not even, %d\n", totalval, totalval % 2)
    return false
  }
  
  eachval := totalval/2
  zeroix := -1
  for ix :=0; ix < len(arr)-1; ix++ {
    eachval -= arr[ix]
    if eachval <= 0 {
      zeroix = ix
      break
    }
  }
  
  if eachval < 0 {
    fmt.Println("can't find mid point")
    return false
  }
  
  if arr[zeroix] < arr[zeroix+1] {
    return true
  }
  
  fmt.Println("all fail")
  return false
}

func main() {
  arr := []int{1, 5, 7, 1}
  fmt.Printf("Input: %+v, is balanced split? %+v\n", arr, balancedSplitExists(arr))
  
    arr = []int{12, 7, 1, 6, 5, 7, 6, 2, 6}
  fmt.Printf("Input: %+v, is balanced split? %+v\n", arr, balancedSplitExists(arr))
  
    arr = []int{1, 5, 7, 1, 3, 3}
  fmt.Printf("Input: %+v, is balanced split? %+v\n", arr, balancedSplitExists(arr))
  
    arr = []int{1, 5, 7, 1, 2, 55, 6, 7, 8, 9, 43, 32, 44, 38, 7, 4, 1}
  fmt.Printf("Input: %+v, is balanced split? %+v\n", arr, balancedSplitExists(arr))
}


