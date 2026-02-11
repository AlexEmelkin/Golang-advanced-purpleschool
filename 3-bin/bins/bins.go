package bins

import (
	"encoding/json"
	"errors"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func (binList *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binList)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	switch {
	case id == "":
		return nil, errors.New("Invalid_Id")
	case name == "":
		return nil, errors.New("Invalid_name")
	default:
		return &Bin{
			Id:        id,
			Private:   private,
			CreatedAt: createdAt,
			Name:      name,
		}, nil

	}
}

func NewBinList(arrB []Bin) *BinList {
	bl := &BinList{}
	bl.Bins = arrB
	return bl

}
