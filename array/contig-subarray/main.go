// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func countSubarrays(arr []int) []int {
  resp := make([]int, len(arr))
  
  for ix := 0; ix < len(arr); ix++ {
    iret := findSub(arr, ix)
    
    resp[ix] = len(iret)
  } 
	return resp;
}

func main() {
  arr := []int{3, 4, 1, 6, 2}
  fmt.Printf("Ret %+v, for input array %+v\n", countSubarrays(arr), arr)
  
               arr = []int{1, 1, 3}
  fmt.Printf("Ret %+v, for input array %+v\n", countSubarrays(arr), arr)

               arr = []int{3, 4, 5, 5, 5, 6, 6, 7, 7, 8}
  fmt.Printf("Ret %+v, for input array %+v\n", countSubarrays(arr), arr)
  
               arr = []int{6,6,6,6,6,6,6,6,6,6,6}
  fmt.Printf("Ret %+v, for input array %+v\n", countSubarrays(arr), arr)
}


func findSub(a []int, ix int) [][]int {
  resp := [][]int{}
  
  // append self
  resp = append(resp, []int{a[ix]})

  allsub := [][]int{}
    
  // first, go left
    
    for x := 1;; x++ {
      if ix - x < 0 {
        break
      }
      if a[ix - x] < a[ix] {
        allsub = append(allsub, a[ix -x: ix])
      } else {
        break
      }
    }
    resp = append(resp, allsub...)
    allsub = [][]int{}
  
  // Now the right
  if ix < len(a) - 1 {
    for x := 1;; x++ {
      if ix + x >= len(a) {
        break
      }
      if a[ix + x] < a[ix] {
        allsub = append(allsub, a[ix: ix+x])
      } else {
        break
      }
    }
    resp = append(resp, allsub...)
  }
  
  return resp
}

