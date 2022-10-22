package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Println("Введите длинну стороны a:")
	fmt.Scanf("%d\n", &a)

	fmt.Println("Введите длинну стороны b:")
	fmt.Scanf("%d\n", &b)

	fmt.Printf("Площадь прямоугольника ровна %d", a*b)
}
