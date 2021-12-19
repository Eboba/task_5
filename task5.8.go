package main

import (
	"fmt"
)

func main() {

	fmt.Println("Игра «Угадай число».")

	fmt.Println("Загадайте число от 1 до 10:")
	fmt.Println("Отвечай на вопрос больше/меньше/угадал.")
	var answer string

	fmt.Println("Ваше число 5?")
	fmt.Scan(&answer)

	if answer == "больше" {
		fmt.Println("Ваше число 7?")
		fmt.Scan(&answer)

		if answer == "больше" {
			fmt.Println("Ваше число 9?")
			fmt.Scan(&answer)

			if answer == "больше" {
				fmt.Println("Ваше число 10!")

			} else if answer == "меньше" {
				fmt.Println("Ваше число 8!")
			} else if answer == "да" {
				fmt.Println("Ура, сыграем еще?")
			}

		} else if answer == "да" {
			fmt.Println("Ура, сыграем еще?")

		} else if answer == "меньше" {
			fmt.Println("Ваше число 6!")

		} else if answer == "да" {
			fmt.Println("Ура, сыграем еще?")
		}

	} else if answer == "меньше" {
		fmt.Println("Ваше число 3?")
		fmt.Scan(&answer)

		if answer == "больше" {
			fmt.Println("Ваше число 4!")
			fmt.Scan(&answer)

		} else if answer == "меньше" {
			fmt.Println("Ваше число 2?")
			fmt.Scan(&answer)
		}

		if answer == "меньше" {
			fmt.Println("Ваше число 1!")

		} else if answer == "да" {
			fmt.Println("Ура, сыграем еще?")
		}

	} else if answer == "да" {
		fmt.Println("Ура, сыграем еще?")
	}
}
