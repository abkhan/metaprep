// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"
import "sort"

func countDistinctTriangles(arr [][]int) int {

  res := map[string]int{}
  for _, tri := range arr {
    sort.Ints(tri)
    key := fmt.Sprintf("%d-%d-%d", tri[0], tri[1], tri[2])
    res[key] = 1
  }
  
  return len(res)
}

func main() {
  
  arr := [][]int{{2, 2, 3}, {3, 2, 2}, {2, 5, 6}}
  fmt.Printf("Input: %+v, res: %d\n", arr, countDistinctTriangles(arr))

  arr = [][]int{{8, 4, 6}, {100, 101, 102}, {84, 93, 173}}
  fmt.Printf("Input: %+v, res: %d\n", arr, countDistinctTriangles(arr))

   arr = [][]int{{2, 2, 3}, {3, 2, 2}, {2, 5, 6}, {6, 2, 5}}
  fmt.Printf("Input: %+v, res: %d\n", arr, countDistinctTriangles(arr))

}

