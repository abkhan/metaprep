package main

import "fmt"

func maxCandies(arr []int, k int) int {
  if len(arr) < k {
    return -1
  }

  total := 0
  for ix := 0; ix < k; ix++ {
    maxix := findHighest(arr)
    if maxix == -1 {
      return total
    }
    
    total += arr[maxix]
    arr[maxix] = arr[maxix]/2
  }
  return total
}

func main() {
  arr := []int{2, 1, 7, 4, 2}
  fmt.Printf("For input: %+v, max choco to eat for 3 >> %d\n", arr, maxCandies(arr, 3))

  arr = []int{2, 1, 7, 4, 2}
  fmt.Printf("For input: %+v, max choco to eat for 4 >> %d\n", arr, maxCandies(arr, 4))
  
    arr = []int{13, 4, 17, 9, 7, 5, 13, 21, 5, 6, 13}
  fmt.Printf("For input: %+v, max choco to eat for 5 >> %d\n", arr, maxCandies(arr, 5))
  
    arr = []int{13, 4, 17, 9, 7, 5, 13, 21, 5, 6, 13}
  fmt.Printf("For input: %+v, max choco to eat for 6 >> %d\n", arr, maxCandies(arr, 6))
}

// improve by using heaps (max @ center heap) and update head when we reduce the count
func findHighest(a []int) int {
  if len(a) < 1 {
    return 0
  }
  
  max := a[0]
  maxindex := 0
  for ix := 0; ix < len(a); ix++ {
    if a[ix] > max {
      max = a[ix]
      maxindex = ix
    }
  }
  return maxindex
}

