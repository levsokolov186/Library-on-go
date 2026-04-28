package main

import (
	"Library/Authentication"
	"fmt"
	"log"
)

func main() {

	fmt.Println("=== Добавляем первого пользователя ===")
	user1 := &Authentication.User{}
	user1.AddUser("Иван Петров", "male", "ivan@example.com", "secure123", 25)

	err := Authentication.AddUserToFile(*user1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("✓ Пользователь 1 добавлен")
	}

	fmt.Println("\n=== Пытаемся добавить пустого пользователя ===")
	emptyUser := Authentication.User{}
	err = Authentication.AddUserToFile(emptyUser)
	if err != nil {
		fmt.Println("✗ Ошибка:", err) // Сработает защита
	}
}
