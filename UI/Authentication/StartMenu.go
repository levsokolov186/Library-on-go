package Authentication

import (
	"fmt"
)

func Entrance() {
	fmt.Println("Добро пожаловать в библиотеку!")
	fmt.Println("Здесь лучшие книги!")
	fmt.Println("Теперь выберите: Войти или Зарегистрироваться")

	var loginOrRegister string
	_, err := fmt.Scanln(&loginOrRegister)
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	if loginOrRegister == "Зарегистрироваться" || loginOrRegister == "зарегистрироваться" {
		fmt.Println("Вы выбрали зарегистрироваться")
		Registration()
		return
	}

	if loginOrRegister == "Войти" || loginOrRegister == "войти" {
		fmt.Println("Вы выбрали войти")
		if err := Login(); err != nil {
			fmt.Println("Ошибка входа:", err)
		}
		return
	}

	fmt.Println("Неверный ввод. Пожалуйста, введите Войти или Зарегистрироваться")
}
