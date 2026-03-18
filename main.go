package main

import (
	"demo/app-4/account"
	"demo/app-4/files"
	"demo/app-4/output"
	"fmt"
)

func main() {
	dialogue()
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин: "})
	password := promptData([]string{"Введите пароль: "})
	url := promptData([]string{"Введите url: "})

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
		choice := promptData([]string{"1. Создать аккаунт", "2. Найти аккаунт", "3. Удалить аккаунт", "4. Выход", "Выберите вариант"})
		switch choice {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func findAccount(vault *account.VaultWithDb) {
	choice := promptData([]string{"Введите url, который хотите найти: "})
	accs := vault.FindAccountsByURL(choice)
	if len(accs) == 0 {
		fmt.Println("Нет аккаунтов с таким url")
	} else {
		for _, val := range accs {
			fmt.Printf("Логин: %s Пароль: %s Url: %s\n", val.Login, val.Password, val.Url)
		}
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	choice := promptData([]string{"Введите url, который хотите удалить: "})
	vault.DeleteAccountsByURL(choice)
	fmt.Println("Аккаунты удалены")
}

func promptData[T any](prompt []T) string {
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
