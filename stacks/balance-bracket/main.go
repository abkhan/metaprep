// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "sync"
import "fmt"
import "errors"

func isBalanced(s string) bool {
  stack := NewStack()
  for ix := 0; ix < len(s); ix++ {
    if s[ix] == '{' || s[ix] == '[' || s[ix] == '(' {
      stack.Push(s[ix])
    }
    if s[ix] == '}' || s[ix] == ']' || s[ix] == ')' {
      st, err := stack.Pop()
      if err != nil {
        return false
      }
      if !match(st, s[ix]) {
        return false
      }
    }
  }
  if len(stack.s) > 0 {
    return false
  }
  return true
}

func main() {
  s := "{[()]}"
  fmt.Printf("is [%s] balanced? %+v\n", s, isBalanced(s))
  
    s = "{}()"
  fmt.Printf("is [%s] balanced? %+v\n", s, isBalanced(s))
  
    s = "{(})"
  fmt.Printf("is [%s] balanced? %+v\n", s, isBalanced(s))
  
    s = ")"
  fmt.Printf("is [%s] balanced? %+v\n", s, isBalanced(s))
  
  s = "{[()]}(){[](}"
  fmt.Printf("is [%s] balanced? %+v\n", s, isBalanced(s))
}



func match(a, b byte) bool {
  if a == '(' && b == ')' {
    return true
  }
    if a == '[' && b == ']' {
    return true
  }
  if a == '{' && b == '}' {
    return true
  }
  return false
}


type stack struct {
     lock sync.Mutex // you don't have to do this if you don't want thread safety
     s []byte
}

func NewStack() *stack {
    return &stack {sync.Mutex{}, make([]byte,0), }
}

func (s *stack) Push(v byte) {
    s.lock.Lock()
    defer s.lock.Unlock()

    s.s = append(s.s, v)
}

func (s *stack) Pop() (byte, error) {
    s.lock.Lock()
    defer s.lock.Unlock()


    l := len(s.s)
    if l == 0 {
        return 0, errors.New("Empty Stack")
    }

    res := s.s[l-1]
    s.s = s.s[:l-1]
    return res, nil
}

