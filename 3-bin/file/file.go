package file

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	extension := filepath.Ext(path)
	if extension != ".json" {
		return nil, errors.New("NOT_JSON_EXT")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func WriteFile() {}
