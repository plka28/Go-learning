package account

import (
	"demo/app-4/encrypter"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts" vault:"accounts"`
	UpdatedAt time.Time `json:"updatedAt" vault:"accounts"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.vault")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	vault.db.Write(encData)
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	result := make([]Account, 0)
	for _, val := range vault.Accounts {
		if checker(val, str) == true {
			result = append(result, val)
		}
	}
	return result
}

func (vault *VaultWithDb) DeleteAccountsByURL(url string) {
	var accounts []Account
	for _, val := range vault.Accounts {
		if strings.Compare(val.Url, url) != 0 {
			accounts = append(accounts, val)
		}
	}
	vault.Accounts = accounts
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	vault.db.Write(encData)
}
