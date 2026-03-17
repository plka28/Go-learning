package main

import (
	"demo/app-4/account"
	"fmt"
)

func main() {
	dialogue()
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL/Неверный логин")
		return
	}

	vault = account.NewVault()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

func dialogue() {
	var choice int
	vault := account.NewVault()
Menu:
	for {
		fmt.Println("1. Создать аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выход")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func findAccount(vault *account.Vault) {
	choice := promptData("Введите url, который хотите найти: ")
	accs := vault.FindAccountsByURL(choice)
	if len(accs) == 0 {
		fmt.Println("Нет аккаунтов с таким url")
	} else {
		for _, val := range accs {
			fmt.Printf("Логин: %s Пароль: %s Url: %s\n", val.Login, val.Password, val.Url)
		}
	}
}

func deleteAccount(vault *account.Vault) {
	choice := promptData("Введите url, который хотите удалить: ")
	vault.DeleteAccountsByURL(choice)
	fmt.Println("Аккаунты удалены")
}
