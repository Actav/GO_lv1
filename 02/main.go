package main

import (
	"fmt"
	"math"
)

func main() {
  var (
    n int
    d []int
  )

  fmt.Print("Введите число: ")
	fmt.Scanln(&n)

  endN := math.Sqrt(float64(n)) + 1
  arr := make([]bool, n)
  for i := 2; i <= int(endN); i++ {
    if arr[i] {
      continue
    }

    for j := i * i; j < n; j += i {
      arr[j] = true
    }
  }


  for k,v := range arr {
    if !v && k > 1 {
      d = append(d, k)
    }
  }

  fmt.Printf("Простые числа от 0 до %d\n %+v", n, d)
}
