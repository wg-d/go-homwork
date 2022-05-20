package main

import "fmt"

func zeroval(ival int) {
  ival = 100 
}

func zeroptr(iptr * int) {
   *iptr = 200
}

func main () {
  i := 1
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)

  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)

}
