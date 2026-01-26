package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор")
	operation, numberStr := getUserInput()
	num2calc, err := str2num(numberStr)
	if err != nil {
		fmt.Println("Ошибка конвертации чисел", err.Error())
		return
	}
	res := calculate(operation, num2calc)
	fmt.Println(res)

}

func getUserInput() (string, string) {
	var operation string
	var numberStr string

	for {
		fmt.Println("Введите операцию(AVG - среднее, SUM - сумму, MED - медиану)")
		fmt.Scan(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			break
		}
		fmt.Println(operation, " - не явлвяется допустимой операцией! Введите заново!")

	}
	fmt.Println("Введите числа через запятую")
	fmt.Scan(&numberStr)
	return operation, numberStr
}

func str2num(numberStr string) ([]int, error) {
	var arr []int
	var str string
	//var num int
	i := 0
	numStr := strings.ReplaceAll(numberStr, " ", "")
	for i < len(numStr) {
		if i == len(numStr)-1 && string(numStr[i]) != "," {
			str = str + string(numStr[i])
		}
		if string(numStr[i]) == "," || i == len(numStr)-1 {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			arr = append(arr, num)
			str = ""
		} else {
			str = str + string(numStr[i])
		}
		i = i + 1
	}
	return arr, nil

}

func calculate(operation string, num2calc []int) float64 {

	var res float64
	if operation == "SUM" {
		for _, n := range num2calc {
			res = res + float64(n)
		}
	} else if operation == "AVG" {
		cnt := 0
		for _, n := range num2calc {
			res = res + float64(n)
			cnt = cnt + 1
		}
		res = res / float64(cnt)
	} else if operation == "MED" {
		res = calculateMedian(num2calc)
	}

	return res
}

func calculateMedian(data []int) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	// 1. Сортируем срез
	sort.Ints(data)

	m := n / 2
	// 2. Если четное количество, берем среднее двух центральных [5]
	if n%2 == 0 {
		return float64((data[m-1] + data[m])) / 2.0
	}
	// 3. Если нечетное, берем центральный элемент [3]
	return float64(data[m])
}
