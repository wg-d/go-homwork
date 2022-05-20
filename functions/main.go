package main

import "fmt"

func plus(a , b int) int{
  return a + b
}

func plusplus(a, b, c int) int {
  return a + b + c
}

func main() {
  res := plus(1,2)
  fmt.Println("1+2=", res)

  res = plusplus(1, 2, 3)
  fmt.Println("1+2+3=", res)

  fmt.Println("1 * 3 = ", times(1,3) )
}

func times(a,b int) int{
  return a * b
}

