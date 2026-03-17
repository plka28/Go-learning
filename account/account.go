package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#*-?")

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (myAccount Account) OutputPassword() {
	color.Cyan(myAccount.Login)
	fmt.Println(myAccount.Password, myAccount.Url)
}

func (acc *Account) generatePassword(n int) {
	password := make([]rune, n)
	for i := range password {
		password[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(password)
}

func NewAccount(login, password, urlstring string) (*Account, error) {
	if login == "" {
		return nil, errors.New("InvalidLogin")
	}

	_, err := url.ParseRequestURI(urlstring)
	if err != nil {
		return nil, errors.New("InvalidURL")
	}
	acc := &Account{
		Login:    login,
		Password: password,
		Url:      urlstring,
	}
	if password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}
