package main

import (
	"HomeTask/3-struct/bins"
	"HomeTask/3-struct/file"
	"HomeTask/3-struct/storage"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {

	binList := bins.NewBinList(storage.NewStorageDb("storage.json"))
Menu:
	for {
		variant := promptData([]string{"1. Создать бинлист", "2. Прочитать  бинлист", "3. Прочиать файл json", "4. Выход", "Выберите вариант"})
		switch variant {
		case "1":
			createBinList(binList)

		case "2":
			readBinList(binList)

		case "3":
			readFile()

		default:
			break Menu
		}
	}

}

func promptData[T any](prompt []T) T {
	var res T
	i := 0
	for i < len(prompt)-1 {
		fmt.Println(prompt[i])
		i = i + 1
	}
	fmt.Print(prompt[i], ":")
	fmt.Scanln(&res)

	return res
}

func createBinList(binList *bins.BinListWithDb) {
	var private bool
	id := promptData([]string{"Введите ИД"})
	privateStr := promptData([]string{"Приватный? y/n"})
	name := promptData([]string{"Введите имя"})
	createdAt := time.Now()
	if privateStr == "y" {
		private = true
	} else {
		private = false
	}
	myBinList, err := bins.NewBin(id, private, createdAt, name)
	if err != nil {
		color.Red("Неверный формат")
		fmt.Println(err)
		return
	}
	binList.AddBin(*myBinList)
}

func readBinList(binList *bins.BinListWithDb) {
	if len(binList.Bins) > 0 {
		fmt.Println("Список bin:")
		for _, bin := range binList.Bins {
			bin.Output()
		}
	} else {
		color.Red("Список bin пуст")
	}
}

func readFile() {
	filepath := promptData([]string{"Введите путь до файла"})
	db := file.NewJsonDb(filepath)
	data, err := db.Read()
	if err != nil {
		color.Red("Ошибка чтения файла", err)
	} else {
		color.Cyan(string(data))
	}

}
