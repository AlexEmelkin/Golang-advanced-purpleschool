package storage

import (
	"fmt"
	"os"
)

type StorageDb struct {
	name string
}

func NewStorageDb(fileName string) *StorageDb {
	return &StorageDb{
		name: fileName,
	}
}

func (db *StorageDb) Save(content []byte) {
	file, err := os.Create(db.name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Запись успешна")

}

func (db *StorageDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.name)
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
