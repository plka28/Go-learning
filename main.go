package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
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

func (myAccount account) outputPassword() {
	fmt.Println(myAccount.login, myAccount.password, myAccount.url)
}

func (acc *account) generatePassword(n int) {
	password := make([]rune, n)
	for i := range password {
		password[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(password)
}

func newAccount(login, password, urlstring string) (*account, error) {
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

func main() {

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL/Неверный логин")
		return
	}

	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
