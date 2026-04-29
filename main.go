package main

import (
	"Library/Authentication"
	"fmt"
	"log"
)

func main() {

	user1 := &Authentication.User{}
	user1.AddUser("Иван Петров", "male", "h@gmail.com", "secure123", 25)
	err := Authentication.AddUserToFile(*user1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("✓ Пользователь 1 добавлен")
	}

	emptyUser := Authentication.User{}
	err = Authentication.AddUserToFile(emptyUser)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("✓ Пользователь 1 добавлен")
	}

}
