package main

import (
	"fmt"
)

func main() {

	fmt.Println("Сумма без сдачи")

	fmt.Println("Введите стоимость товара:")
	var price int
	fmt.Scan(&price)

	fmt.Println("Введите номинал первой монеты:")
	var one int
	fmt.Scan(&one)

	fmt.Println("Введите номинал второй монеты:")
	var two int
	fmt.Scan(&two)

	fmt.Println("Введите номинал третьей монеты:")
	var three int
	fmt.Scan(&three)

	if one < 0 || two < 0 || three < 0 || price <= 0 {
		fmt.Println("Значение введено не верно.")
	} else if one+two+three == 0 {
		fmt.Println("Денег нет, но Вы держитесь(с).")
	} else if one == price || two == price || three == price {
		fmt.Println("За товар можно заплатить без сдачи.")
	} else if one+two == price || two+three == price || three+one == price {
		fmt.Println("За товар можно заплатить без сдачи.")
	} else if one+two+three == price {
		fmt.Println("За товар можно заплатить без сдачи.")
	} else {
		fmt.Println("За товар нельзя заплатить без сдачи.")
	}
}
