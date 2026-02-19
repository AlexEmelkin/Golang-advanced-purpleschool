package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type JsonDb struct {
	filepath string
}

func NewJsonDb(path string) *JsonDb {
	return &JsonDb{
		filepath: path,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	extension := filepath.Ext(db.filepath)
	if extension != ".json" {
		return nil, errors.New("NOT_JSON_EXT")
	}
	data, err := os.ReadFile(db.filepath)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filepath)
	if err != nil {
		//output.PrintError(err)
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		//output.PrintError(err)
		fmt.Println("WriteString", err)
		return
	}
	fmt.Println("Запись успешна")
}
