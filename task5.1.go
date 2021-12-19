package main

import (
	"fmt"
)

func main() {

	fmt.Println("Определение координатной плоскости точки.")

	fmt.Println("Введите x:")
	var x int
	fmt.Scan(&x)

	fmt.Println("Введите y:")
	var y int
	fmt.Scan(&y)

	if x == 0 && y == 0 {
		fmt.Println("Точка находится в 0 координате.")
	} else if x == 0 && y != 0 {
		fmt.Println("Точка находится на оси у.")
	} else if x != 0 && y == 0 {
		fmt.Println("Точка находится на оси x.")
	} else if x > 0 && y > 0 {
		fmt.Println("Точка принадлежит 1 координатной четверти.")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка принадлежит 2 координатной четверти.")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка принадлежит 3 координатной четверти.")
	} else {
		fmt.Println("Точка принадлежит 4 координатной четверти.")
	}
}
