// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"
import "errors"

func findMaxProduct(arr []int) []int {
  retval := []int{}
  for ix := 0; ix < len(arr); ix++ {
    if ix < 2 {
      retval = append(retval, -1)
      continue
    }

    reta, err := find3x(arr[:ix+1])
    if err != nil {
      retval = append(retval, -1)
    } else {
      retval = append(retval, reta[0] * reta[1] * reta[2])
  }
}
  return retval
}

func main() {
  arr := []int{1,2,3,4,5}
  fmt.Printf("Input> %+v, >>> retval >>> %+v\n", arr, findMaxProduct(arr))

  arr = []int{2, 1, 2, 1, 2}
  fmt.Printf("Input> %+v, >>> retval >>> %+v\n", arr, findMaxProduct(arr))
  
    arr = []int{4,5,6, 12, 13, 14, 2, 3, 4, 1, 1, 2, 15}
  fmt.Printf("Input> %+v, >>> retval >>> %+v\n", arr, findMaxProduct(arr))
}


func find3x(a []int) ([3]int, error) {
  reta := [3]int{}

  if len(a) < 3 {
    return reta, errors.New("less than 3") 
  }
  reta[0] = a[0]
  reta[1] = a[1]
  reta[2] = a[2]
  
  if len(a) == 3 {
    return reta, nil
  }
  
  for ix := 3; ix < len(a); ix++ {
    for ri := 0; ri < 3; ri++ {
      if a[ix] > reta[ri] {
        reta[ri] = a[ix]
        break
      }
    }
  }
  
  return reta, nil
}

