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

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (myAccount account) OutputPassword() {
	color.Cyan(myAccount.login)
	fmt.Println(myAccount.password, myAccount.url)
}

func (acc *account) generatePassword(n int) {
	password := make([]rune, n)
	for i := range password {
		password[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(password)
}

func NewAccount(login, password, urlstring string) (*account, error) {
	if login == "" {
		return nil, errors.New("InvalidLogin")
	}

	_, err := url.ParseRequestURI(urlstring)
	if err != nil {
		return nil, errors.New("InvalidURL")
	}
	acc := &account{
		login:    login,
		password: password,
		url:      urlstring,
	}
	if password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}
