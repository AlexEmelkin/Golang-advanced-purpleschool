package main

import "fmt"

func main() {
	const Usd2Eur = 0.94
	const Usd2Rub = 79.0
	const Eur2Rub = Usd2Rub / Usd2Eur
}

func Converter(cur float64, originalCur string, targetCur string) {

}

func getUserInput() (float64, string, string) {
	var cur float64
	var originalCur string
	var targetCur string
	fmt.Print("ВВедите сумму в исходной валюте: ")
	fmt.Scan(&cur)
	fmt.Print("ВВедите исходную валюту ")
	fmt.Scan(&originalCur)
	fmt.Print("ВВедите целевую валюту: ")
	fmt.Scan(&targetCur)
	return cur, originalCur, targetCur
}
