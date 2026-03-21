package api

import (
	"HomeTask/3-struct/bins"
	"HomeTask/3-struct/config"
	"HomeTask/3-struct/file"
	"HomeTask/3-struct/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type MetaData struct {
	MetaData bins.Bin `json:"metadata"`
}

var conf = GetConfig()
var urlBin = "https://api.jsonbin.io/v3/b"
var binList = bins.NewBinList(storage.NewStorageDb("list.json"))

func GetConfig() *config.Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	conf := *config.NewConfig()
	conf.Key = "$2a$10$DijyuNp8Aq/Zit7XtipBUONtqbdHK9JKGscqqCBJiBz2jGYxijwQy"
	return &conf
}

func CreateBin(fileName string, binName string) (error, string) {
	db := file.NewJsonDb(fileName)
	postBody, err := db.Read()
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	req, err := http.NewRequest("POST", urlBin, bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", conf.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}
	//fmt.Println(string(body))
	var mD MetaData
	json.Unmarshal(body, &mD)
	if mD.MetaData.Name == "" {
		mD.MetaData.Name = binName
	}

	binList.AddBin(mD.MetaData)
	fmt.Printf("Cretae %s success", binName)

	return nil, mD.MetaData.Id
}

func GetBin(binId string) string {
	baseUrl, err := url.Parse(urlBin + "/" + binId)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", conf.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)
}

func UpdateBin(fileName string, binId string) error {

	db := file.NewJsonDb(fileName)
	postBody, err := db.Read()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	baseUrl, err := url.Parse(urlBin + "/" + binId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	req, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", conf.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Update %s success", binId)

	return nil

}

func DeleteBin(binId string) error {
	baseUrl, err := url.Parse(urlBin + "/" + binId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	req, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", conf.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Delete %s success", binId)

	return nil
}

func ReadList() {
	if len(binList.Bins) > 0 {
		fmt.Println("Список bin:")
		for _, bin := range binList.Bins {
			bin.Output()
		}
	} else {
		color.Red("Список bin пуст")
	}
}
