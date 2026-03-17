package account

import (
	"demo/app-4/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccountsByURL(url string) []Account {
	result := make([]Account, 0)
	for _, val := range vault.Accounts {
		if strings.Compare(val.Url, url) == 0 {
			result = append(result, val)
		}
	}
	return result
}

func (vault *Vault) DeleteAccountsByURL(url string) {
	var accounts []Account
	for _, val := range vault.Accounts {
		if strings.Compare(val.Url, url) != 0 {
			accounts = append(accounts, val)
		}
	}
	vault.Accounts = accounts
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}
