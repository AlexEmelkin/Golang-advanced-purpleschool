package bins

import (
	"HomeTask/3-struct/storage"
	"encoding/json"
	"errors"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

type BinListWithDb struct {
	BinList
	db Db
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

/*func NewBinList(arrB []Bin) *BinList {
	bl := &BinList{}
	bl.Bins = arrB
	return bl

}*/

func NewBinList(db Db) *BinListWithDb {
	//db := files.NewJsonDb("data.json")
	file, err := db.Read()
	if err != nil {
		return &BinListWithDb{
			BinList: BinList{
				Bins: []Bin{},
			},
			db: db,
		}
	}

	var binList BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		//("Не удалось разобрать файл " + err.Error())
		color.Red("Не удалось разобрать файл " + err.Error())
		return &BinListWithDb{
			BinList: BinList{
				Bins: []Bin{},
			},
			db: db,
		}
	}
	return &BinListWithDb{
		BinList: binList,
		db:      db,
	}
}

func (binList *BinListWithDb) AddBin(bin Bin) {

	binList.Bins = append(binList.Bins, bin)
	data, err := binList.BinList.ToBytes()
	if err != nil {
		//output.PrintError(err)
		color.Red("Не удалось преобразовать файл " + err.Error())
	}
	storage.SaveBinList(data, "binList.json")
}
