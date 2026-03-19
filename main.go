package main

import (
	"demo/app-4/account"
	"demo/app-4/files"
	"demo/app-4/output"
	"fmt"
	"strings"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount, "2": findAccountByUrl, "3": findAccountByLogin, "4": deleteAccount,
}

func main() {
	dialogue()
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL/Неверный логин")
		return
	}

	vault = account.NewVault(files.NewJsonDb("data.json"))
	vault.AddAccount(*myAccount)
}

func dialogue() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		choice := promptData("1. Создать аккаунт", "2. Найти аккаунт по url", "3. Найти аккаунт по логину", "4. Удалить аккаунт", "5. Выход", "Выберите вариант")
		menuFunc := menu[choice]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		// switch choice {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {
	choice := promptData("Введите url, который хотите найти: ")
	accs := vault.FindAccounts(choice, func(a account.Account, s string) bool { return strings.Contains(a.Url, s) })
	if len(accs) == 0 {
		fmt.Println("Нет аккаунтов с таким url")
	} else {
		for _, val := range accs {
			fmt.Printf("Логин: %s Пароль: %s Url: %s\n", val.Login, val.Password, val.Url)
		}
	}
}
func findAccountByLogin(vault *account.VaultWithDb) {
	choice := promptData("Введите url, который хотите найти: ")
	accs := vault.FindAccounts(choice, func(a account.Account, s string) bool { return strings.Contains(a.Login, s) })
	if len(accs) == 0 {
		fmt.Println("Нет аккаунтов с таким login")
	} else {
		for _, val := range accs {
			fmt.Printf("Логин: %s Пароль: %s Url: %s\n", val.Login, val.Password, val.Url)
		}
	}
}
func deleteAccount(vault *account.VaultWithDb) {
	choice := promptData("Введите url, который хотите удалить: ")
	vault.DeleteAccountsByURL(choice)
	fmt.Println("Аккаунты удалены")
}

func promptData(prompt ...any) string {
	for index, value := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v: ", value)
		} else {
			fmt.Println(value)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func factory() func(*account.VaultWithDb) {
	return findAccountByUrl
}
