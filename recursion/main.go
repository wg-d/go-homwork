package main

import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n - 1)
}

func secondFib(n int) int {
  if n <2 {
    return n
  }

  return secondFib(n - 1) + secondFib(n - 2)

}

func main() {
  fmt.Println(fact(7))

  var fib func(n int) int

  fib = func(n int) int {
    if n <2 {
      return n
    }
    return fib(n - 1) + fib(n - 2)
  }

  fmt.Println(fib(7))
  fmt.Println(secondFib(7))
}
