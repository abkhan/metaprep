package main

import "fmt"

func findEncryptedWord(s string) string {
  if len(s) < 3 {
    return s
  }
  if len(s) == 3 {
    return string(s[1]) + string(s[0]) + string(s[2])
  }
  entrix := len(s)/2
  if len(s) % 2 == 0 {
    entrix = len(s)/2 -1
  }
  
  return string(s[entrix]) + findEncryptedWord(s[0:entrix]) + findEncryptedWord(s[entrix+1:])
  
}

func main() {
  s := "abc"
  fmt.Printf("for [%s], encrypted >> %s\n", s, findEncryptedWord(s))
  
  s = "abcd"
  fmt.Printf("for [%s], encrypted >> %s\n", s, findEncryptedWord(s))

    s = "abcxcba"
  fmt.Printf("for [%s], encrypted >> %s\n", s, findEncryptedWord(s))

    s = "facebook"
  fmt.Printf("for [%s], encrypted >> %s\n", s, findEncryptedWord(s))

}

