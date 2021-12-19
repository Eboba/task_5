package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Решение квадратного уравнения.")

	fmt.Println("Введите коэффициент a:")
	var a float64
	fmt.Scan(&a)

	fmt.Println("Введите коэффициент b:")
	var b float64
	fmt.Scan(&b)

	fmt.Println("Введите коэффициент c:")
	var c float64
	fmt.Scan(&c)

	D := (b*b - 4*a*c)

	if D < 0 {
		fmt.Println("Корней нет.")
	} else if D > 0 {
		fmt.Println((-b+math.Sqrt(D))/(2*a), (-b-math.Sqrt(D))/(2*a))
	} else {
		fmt.Println((-b + math.Sqrt(D)) / (2 * a))
	}
}
