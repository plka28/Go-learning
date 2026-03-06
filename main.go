package main

import "fmt"

func main() {
	bookmark := map[string]string{}
	for {
		switch dialogue() {
		case 1:
			show(bookmark)
		case 2:
			bookmark = add(bookmark)
		case 3:
			bookmark = del(bookmark)
		case 4:
			return
		}
	}
}
func dialogue() int {
	for {
		var choice int
		fmt.Println("1. Посмотреть закладки")
		fmt.Println("2. Добавить закладку")
		fmt.Println("3. Удалить закладку")
		fmt.Println("4. Выход")
		fmt.Scan(&choice)
		if choice == 1 || choice == 2 || choice == 3 || choice == 4 {
			return choice
		}
		fmt.Println("Неправильный ввод")
	}
}
func show(bookmark map[string]string) {
	for key, value := range bookmark {
		fmt.Printf("Ключ: %s, Значение: %s\n", key, value)
	}
}
func add(bookmark map[string]string) map[string]string {
	var newKey string
	var newValue string
	fmt.Print("Введите ключ: ")
	fmt.Scan(&newKey)
	fmt.Print("Введите значение: ")
	fmt.Scan(&newValue)
	bookmark[newKey] = newValue
	return bookmark
}
func del(bookmark map[string]string) map[string]string {
	var newKey string
	fmt.Print("Введите ключ: ")
	fmt.Scan(&newKey)
	delete(bookmark, newKey)
	return bookmark
}
