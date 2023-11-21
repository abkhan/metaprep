// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func main() {
  tsa := "abcd"
  tsb := "abcd"
  fmt.Printf("For %s and %s, max: %d\n", tsa, tsb, matchingPairs(tsa, tsb))
  
    tsa = "adcb"
  tsb = "abcd"
  fmt.Printf("For %s and %s, max: %d\n", tsa, tsb, matchingPairs(tsa, tsb))
  
    tsa = "mnopqr"
  tsb = "mropqz"
  fmt.Printf("For %s and %s, max: %d\n", tsa, tsb, matchingPairs(tsa, tsb))
  
    tsa = "xyzabc"
  tsb = "abcxyz"
  fmt.Printf("For %s and %s, max: %d\n", tsa, tsb, matchingPairs(tsa, tsb))
}

func matchingPairs(s string, t string) int {

  smap, tmap := match(s, t)
  if len(smap) == 0 {
    return len(s) -2
  }
  
  // find out how many non-matching chars can be matched
  count := 0
  for _, tc := range tmap {
    if _, yes := smap[tc]; yes {
      count++
    }
  }
  if count > 1 {
    count = 2
  }
  return (len(s) - len(smap)) + count
}


// returns from t, unmatching by char to index
// returns from s, unmatching by index to char
func match(s string, t string) (map[byte]int, map[int]byte) {
  smap := map[byte]int{}
  tmap := map[int]byte{}
  
  for ix := 0; ix < len(s); ix++ {
    if s[ix] != t[ix] {
      smap[s[ix]] = ix
      tmap[ix] = t[ix]
    }
  }
  
  return smap, tmap
}

