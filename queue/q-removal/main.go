package main

import "fmt"

type node struct {
  val int
  pos int
}

func findPositions(arr []int, x int) []int {
  resp := []int{}
  work := make(chan node, len(arr))
  
  for ix := 0; ix < len(arr); ix++ {
    work <- node{
      val: arr[ix], 
      pos: ix+1,
    }
  }
  
  for ix := 0; ix < x; ix++ {
    npop := x
    if len(work) < x {
      npop = len(work)
    }
    
    pq := []node{}
    for ix := 0; ix < npop; ix++ {
      node := <- work
      pq = append(pq, node)
    }
    //fmt.Printf("workQ len: %d\n", len(work))
    
    pq, rem := remove(pq)
    resp = append(resp, rem)
    for ix := 0; ix < len(pq); ix++ {
      work <- pq[ix]
    }
    //fmt.Printf("end - workQ len: %d\n", len(work))
  }
  return resp
}

func main() {
  arr := []int{1, 2, 2, 3, 4, 5}
  x := 5
  fmt.Printf("Q: %+v, where x=%d, output: %+v\n", arr, x, findPositions(arr, x))

  arr = []int{1, 2, 2, 3, 4, 5}
  x = 3
  fmt.Printf("Q: %+v, where x=%d, output: %+v\n", arr, x, findPositions(arr, x))
  
    arr = []int{6, 7, 1, 2, 2, 3, 4, 5, 8}
  x = 6
  fmt.Printf("Q: %+v, where x=%d, output: %+v\n", arr, x, findPositions(arr, x))
}

// remove would find max, or take the first, and return original position
func remove(a []node) ([]node, int) {
  if len(a) < 1 {
    return a, -1
  }
  
  //fmt.Printf("remove called: %+v\n", a)
  
  maxval := 0;
  maxix := -1;
  for ix := 0; ix < len(a); ix++ {
    if a[ix].val > maxval {
      maxval = a[ix].val
      maxix = ix
    }
  }
  
  if maxix == -1 {
    origix := a[0].pos
    a = a[1:]
    return a, origix
  }
  
  origix := a[maxix].pos
  
  // dec values
  for ix := 0; ix < len(a); ix++ {
    if a[ix].val != 0 {
      a[ix].val -= 1
    }
  }
  ra := append(a[:maxix], a[maxix+1:]...)
  //fmt.Printf("after remove: %+v\n", a)
  return ra, origix
}

