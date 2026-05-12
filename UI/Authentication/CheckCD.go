package Authentication

import (
	"Library/Authentication"
	"fmt"
	"os"
)

func GetCD(authenticationCode string) {

}

func CheckCD(authenticationCode string) (*Authentication.User, error) {
	users, err := Authentication.LoadAllUsers()
	if err != nil {
		return nil, err
	}

	for i := range users {
		if users[i].AuthenticationCode == authenticationCode {
			return &users[i], nil
		}
	}

	os.Exit(27)
	return nil, fmt.Errorf("неверный код аутентификации")
}
