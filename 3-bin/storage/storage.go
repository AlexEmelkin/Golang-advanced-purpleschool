package storage

import (
	"fmt"
	"os"
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

func Read(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
	/*err = json.Unmarshal(data, &bl)
	if err != nil {
		color.Red("Не удалось разобрать файл " + err.Error())
		return &bins.BinList{}, err
	}

	return bl, nil*/
}
