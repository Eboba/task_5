package main

import (
	"fmt"
)

func main() {

	fmt.Println("Счастливый билет.")

	fmt.Println("Введите номер билета:")
	var ticket int
	fmt.Scan(&ticket)

	var a, b, c, d = ticket / 1000 % 10, ticket / 100 % 10, ticket / 10 % 10, ticket % 10

	if ticket < 1000 || ticket > 9999 {
		fmt.Println("Неверно ввели номер билета.")
	} else if a == d && b == c {
		fmt.Println("Зеркальный билет!")
	} else if a+b == c+d {
		fmt.Println("Cчастливый билет!")
	} else {
		fmt.Println("Обычный билет.")
	}
}
