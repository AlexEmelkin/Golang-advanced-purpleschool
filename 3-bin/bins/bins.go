package bins

import (
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
