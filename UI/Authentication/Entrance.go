package Authentication

import (
	"Library/Authentication"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AuthenticateUserByEmailAndPassword(email, password string) (*Authentication.User, error) {
	users, err := Authentication.LoadAllUsers()
	if err != nil {
		return nil, err
	}

	for i := range users {
		if users[i].Email == email && users[i].Password == password {
			return &users[i], nil
		}
	}

	return nil, fmt.Errorf("неверный email или пароль")
}

func Login() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Введите пароль: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	user, err := AuthenticateUserByEmailAndPassword(email, password)
	if err != nil {
		return err
	}

	fmt.Printf("Вы успешно вошли, код аундефикаци прилегающий к этому аккаунту: %s\n", user.AuthenticationCode)
	return nil
}
