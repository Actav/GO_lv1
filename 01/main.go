package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

  fmt.Print("Введите число: ")
	fmt.Scanln(&a)

  fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
  fmt.Scanln(&op)

  if op != "!" {
    fmt.Print("Введите второе число: ")
    fmt.Scanln(&b)
  }

	switch op {
  	case "+":
  		{
  			res = a + b
  		}
  	case "-":
  		{
  			res = a - b
  		}
  	case "*":
  		{
  			res = a * b
  		}
  	case "/":
  		{
  			if b == 0 {
  				fmt.Println("\n\nНа ноль делить нельзя")
  				os.Exit(1)
  			}
  			res = a / b
  		}
  	case "^":
  		{
  			res = math.Pow(a, b)
  		}
  	case "!":
  		{
        _, f := math.Modf(a)
        if a < 0 || f != 0 {
  				fmt.Println("\n\nВведенное число должно быть неотрицательным целым числом")
  				os.Exit(1)
  			}

        res = 1.00
        for k := 1.00; k <= a; k++ {
          res *= k
        }
  		}
  	default:
  		{
  			fmt.Println("\n\nОперация выбрана неверно")
  			os.Exit(1)
  		}
	}

  iRes, fRes := math.Modf(res)
  if fRes == 0 {
    fmt.Printf("Результат выполнения операции: %.0f\n", iRes)
  } else {
    fmt.Printf("Результат выполнения операции: %.2f\n", res)
  }
}
