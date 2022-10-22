package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64

	fmt.Println("Введите площадь круга:")
	fmt.Scanf("%f", &s)

	fmt.Printf(
		"Длинна окружности по заданной площади круга ровна %f\n",
		math.Sqrt(s*4*math.Pi),
	)

	fmt.Printf(
		"Диаметр окружности по заданной площади круга равен %f\n",
		2*math.Sqrt(s/math.Pi),
	)
}
