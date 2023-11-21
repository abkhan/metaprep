// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main
import "fmt"

// f : rotation factor
func rotationalCipher(input string, f int) string { 
  resp := ""
  for _, c := range input {
    if isAlphaNumeric(c) {
      if isUpper(c) {
        c = rotateUpper(c, int32(f))
      } else if isLower(c) {
        c = rotateLower(c, int32(f))
      } else if isNum(c) {
        c = rotateNum(c, int32(f))
      }
    }
    resp += string(c)
  }
  return resp
}

func main() {
  ts := "Zebra-493?"
  fmt.Printf("Input: %s >>> %s\n", ts, rotationalCipher(ts, 3))
  ts = "abcdefghijklmNOPQRSTUVWXYZ0123456789"
  fmt.Printf("Input: %s >>> %s\n", ts, rotationalCipher(ts, 39))}


func isAlphaNumeric(c rune) bool {

  if isUpper(c) || isLower(c) || isNum(c) {
    return true
  }
  return false
}

func isUpper(c rune) bool {
    
  if c >= 'A' && c <= 'Z' {
    return true
  }
  return false
}

func isLower(c rune) bool {
    if c >= 'a' && c <= 'z' {
    return true
  }
  return false
}

func isNum(c rune) bool {
    if c >= '0' && c <= '9' {
    return true
  }
  return false
}


func rotateUpper(c rune, f int32) rune {
  return ((c - 'A') + f) % ('Z' - 'A' +1) + 'A'
}


func rotateLower(c rune, f int32) rune {
  return ((c - 'a') + f) % ('z' - 'a' +1) + 'a'
}


func rotateNum(c rune, f int32) rune {
  return ((c - '0') + f) % ('9' - '0' +1) + '0'
}
