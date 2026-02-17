package storage

import (
	"HomeTask/3-struct/bins"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func SaveBinList(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		return err
	}
	fmt.Println("Запись успешна")
	return nil

}

func ReadBinList(name string) (*bins.BinList, error) {
	var bl *bins.BinList
	data, err := os.ReadFile(name)
	if err != nil {
		return &bins.BinList{}, err
	}
	err = json.Unmarshal(data, &bl)
	if err != nil {
		color.Red("Не удалось разобрать файл " + err.Error())
		return &bins.BinList{}, err
	}

	return bl, nil
}
