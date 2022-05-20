package main

import "fmt"

func vals() (int, int) {
  return 3, 7
}

func multipleVals(a, b int) (int, int) {
  return a + b, a * b
} 

func main() {
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)

  sp, _ := multipleVals(5,6)
  fmt.Println(sp)

}
