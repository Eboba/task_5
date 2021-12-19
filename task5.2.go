package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проверить, что одно из чисел — положительное.")

	fmt.Println("Введите первое число:")
	var one int
	fmt.Scan(&one)

	fmt.Println("Введите второе число:")
	var two int
	fmt.Scan(&two)

	fmt.Println("Введите третье число:")
	var three int
	fmt.Scan(&three)

	if one > 0 || two > 0 || three > 0 {
		fmt.Println("Одно число является положительным.")
	} else {
		fmt.Println("Положительных чисел нету.")
	}
}
