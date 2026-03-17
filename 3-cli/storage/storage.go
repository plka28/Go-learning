package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/plka28/3-cli/bins"
	"github.com/plka28/3-cli/file"
)

func BinToJson(bin *bins.Bin, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return errors.New("Не удалось создать файл")
	}
	defer file.Close()
	data, err := json.Marshal(bin)
	if err != nil {
		return errors.New("Не удалось десериализовать Bin")
	}
	_, err = file.Write(data)
	if err != nil {
		return errors.New("Не удалось записать")
	}
	return nil
}

func JsonToBinlist(name string) (bins.BinList, error) {
	binList := make(bins.BinList, 0)
	data, err := file.ReadJsonFile(name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, errors.New("Не удалось десериализовать Bin")
	}
	return binList, nil
}
