package main

import "fmt"

const Usd2Eur = 0.94
const Usd2Rub = 79.0
const Eur2Rub = Usd2Rub / Usd2Eur

func main() {
	fmt.Println("Калькулятор валют")
	cur, original, target := getUserInput()
	result := converter(cur, original, target)
	fmt.Printf("Сумма в новой валюте:%.2f ", result)

}

func converter(cur float64, originalCur string, targetCur string) float64 {
	var res float64
	switch {
	case originalCur == "USD" && targetCur == "EUR":
		res = cur * Usd2Eur
	case originalCur == "EUR" && targetCur == "USD":
		res = cur / Usd2Eur
	case originalCur == "USD" && targetCur == "RUB":
		res = cur * Usd2Rub
	case originalCur == "RUB" && targetCur == "USD":
		res = cur / Usd2Rub
	case originalCur == "EUR" && targetCur == "RUB":
		res = cur * Eur2Rub
	case originalCur == "RUB" && targetCur == "EUR":
		res = cur / Eur2Rub
	}

	return res

}

func getUserInput() (float64, string, string) {
	var curSum float64
	var originalCur string
	var targetCur string
	var targetText string

	for {
		fmt.Print("Введите исходную валюту (USD,EUR,RUB) ")
		fmt.Scan(&originalCur)
		if originalCur == "USD" || originalCur == "EUR" || originalCur == "RUB" {
			break
		}
		fmt.Println(originalCur, " - не явлвяется валютой! Введите заново!")

	}
	for {
		fmt.Print("Введите сумму в исходной валюте: ")
		fmt.Scan(&curSum)
		if curSum > 0 {
			break
		} else {
			fmt.Println("Сумма не может быть отрицательной!")
		}
	}
	for {
		switch originalCur {
		case "USD":
			targetText = "Введите целевую валюту(EUR,RUB): "
		case "EUR":
			targetText = "Введите целевую валюту(USD,RUB): "
		default:
			targetText = "Введите целевую валюту(USD,EUR): "
		}

		fmt.Print(targetText)
		fmt.Scan(&targetCur)
		if (targetCur == "USD" && originalCur != "USD") || (targetCur == "EUR" && originalCur != "EUR") || (targetCur == "RUB" && originalCur != "RUB") {
			break
		} else if (targetCur == "USD" && originalCur == "USD") || (targetCur == "EUR" && originalCur == "EUR") || (targetCur == "RUB" && originalCur == "RUB") {
			fmt.Println("Нельзя выбирать валюту совпадающую с исходной:", targetCur)
		} else {
			fmt.Println(targetCur, " - не явлвяется валютой! Введите заново!")
		}
	}
	return curSum, originalCur, targetCur
}
