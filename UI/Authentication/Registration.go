package Authentication

import (
	"Library/Authentication"
	"fmt"
	"log"
)

//Name               string `json:"name"`
//Age                int    `json:"age"`
//Gender             string `json:"gender"`
//Email              string `json:"email"`
//Password           string `json:"password"`
//AuthenticationCode string `json:"authentication_code"`

func Registration() {
	fmt.Println("Регистрация:")
	fmt.Println("Введите имя, возраст, ваш гендр, имейл, пароль. После заполнения всех данных вы получите код аудефикации и сможете войти в библиотеку.")
	fmt.Println("Пример:")
	fmt.Println("Иван Петров, 25, муж, testUser@gmail.com, secure123, код получите позже")
	var name string
	var age int
	var gender string
	var email string
	var password string

	fmt.Println("Имя: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	fmt.Println("Возвраст: ")
	_, errOne := fmt.Scanln(&age)
	if errOne != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	fmt.Println("Гендер: ")
	_, errTwo := fmt.Scanln(&gender)
	if errTwo != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	fmt.Println("Имейл: ")
	_, errThree := fmt.Scanln(&email)
	if errThree != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	fmt.Println("Пароль: ")
	_, errFour := fmt.Scanln(&password)
	if errFour != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	user := &Authentication.User{}
	user.AddUser(name, gender, email, password, age)
	errFive := Authentication.AddUserToFile(*user)
	if errFive != nil {
		log.Println(err)
	} else {
		fmt.Println("✓ Пользователь добавлен")
	}
}
