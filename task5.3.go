package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проверить, что есть совпадающие числа.")

	fmt.Println("Введите первое число:")
	var one int
	fmt.Scan(&one)

	fmt.Println("Введите второе число:")
	var two int
	fmt.Scan(&two)

	fmt.Println("Введите третье число:")
	var three int
	fmt.Scan(&three)

	if one == two || three == two || three == one {
		fmt.Println("Два числа или более числа совпадают")
	}
}
