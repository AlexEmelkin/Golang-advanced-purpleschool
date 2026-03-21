package api_test

import (
	"HomeTask/3-struct/api"
	"testing"
)

func TestCreateBin(t *testing.T) {

	fileName := "storage.json"
	binName := "testCretae"
	errBinId := ""
	err, binId := api.CreateBin(fileName, binName)
	if err != nil {
		t.Error("Create error:" + err.Error())
		return
	}
	if binId == errBinId {
		t.Error("Empty Bin")
		return
	}

	err = api.DeleteBin(binId)
	if err != nil {
		t.Error("Delete error" + err.Error())
	}

}

func TestUpdateBin(t *testing.T) {

	fileName := "storage.json"
	binName := "testUpdate"
	err, binId := api.CreateBin(fileName, binName)
	if err != nil {
		t.Error("Create error:" + err.Error())
		return
	}
	err = api.UpdateBin(fileName, binId)
	if err != nil {
		t.Error("Update error:" + err.Error())
		return
	}

	err = api.DeleteBin(binId)
	if err != nil {
		t.Error("Delete error" + err.Error())
	}

}

func TestGetBin(t *testing.T) {

	fileName := "storage.json"
	binName := "testCretae"
	err, binId := api.CreateBin(fileName, binName)
	if err != nil {
		t.Error("Create error:" + err.Error())
		return
	}
	str := api.GetBin(binId)
	if str == "" {
		t.Error("Empty Bin")
		return
	}

	err = api.DeleteBin(binId)
	if err != nil {
		t.Error("Delete error" + err.Error())
	}

}

func TestDeleteBin(t *testing.T) {

	fileName := "storage.json"
	binName := "testCretae"
	err, binId := api.CreateBin(fileName, binName)
	if err != nil {
		t.Error("Create error:" + err.Error())
		return
	}

	err = api.DeleteBin(binId)
	if err != nil {
		t.Error("Delete error" + err.Error())
	}

}
