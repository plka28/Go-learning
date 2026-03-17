package file

import (
	"errors"
	"os"
	"strings"
)

func ReadJsonFile(path string) ([]byte, error) {
	if strings.HasSuffix(path, ".json") {
		file, err := os.ReadFile(path)
		if err != nil {
			return nil, errors.New("Не удалось прочитать файл")
		}
		return file, nil
	}
	return nil, errors.New("Не удалось прочитать файл")
}
