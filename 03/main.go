package main

import (
	"fmt"
)

func main() {
	var userNumber int

	fmt.Println("Введите трехзначное число:")
	fmt.Scanf("%d", &userNumber)

	units := userNumber % 10
	tens := int(userNumber/10) % 10
	hundreds := int(userNumber/100) % 10

	fmt.Printf(
		"\nВ числе %d количество сотен: %d, десятков: %d, единиц: %d\n",
		userNumber,
		hundreds,
		tens,
		units,
	)
}
