package main

import "fmt"
import "sort"

func getTotalTime(arr []int) int {
  fmt.Printf("InputArray: %+v\n", arr)
  fmt.Printf(" > unsorted processing result: %+v\n", fromStart(arr, 0))
  sort.Ints(arr)
  fmt.Printf("  >> sorted processing result: %+v\n", fromStart(arr, 0))
  sort.Sort(sort.Reverse(sort.IntSlice(arr)))
  fmt.Printf("  >> descending sorted processing result: %+v\n", fromStart(arr, 0))

  return 0
}

func main() {
  getTotalTime([]int{4, 2, 1, 3})
  getTotalTime([]int{5, 4, 1, 2, 7, 3})
}



func fromStart(a []int, total int) int {
  
  //  fmt.Printf(" -- fromStart InputArray: %+v, total \n", a, total)
  
  if len(a) < 1 {
    return total
  }
  if len(a) == 1 {
    return a[0] + total
  }
  if len(a) == 2 {
    return a[0] + a[1] + total
  }
  ret := fromStart( append([]int{ a[0] + a[1]}, a[2:]...), total + a[0] + a[1])
  return ret
}

