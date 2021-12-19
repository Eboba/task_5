package main

import (
	"fmt"
)

func main() {

	fmt.Println("Определение максимальных процентов.")

	fmt.Println("Введите первую процентную ставку:")
	var one int
	fmt.Scan(&one)

	fmt.Println("Введите второй процентную ставку:")
	var two int
	fmt.Scan(&two)

	fmt.Println("Введите третью процентную ставку:")
	var three int
	fmt.Scan(&three)

	if one == two && two == three {
		fmt.Println("Cтавки равнозначны.")
	} else if (one >= two && two >= three) || (two >= one && one >= three) {
		fmt.Println("Выгодные ставки", one, two)
	} else if (two >= three && three >= one) || (three >= two && two >= one) {
		fmt.Println("Выгодные ставки", two, three)
	} else {
		fmt.Println("Выгодные ставки", one, three)
	}
}
