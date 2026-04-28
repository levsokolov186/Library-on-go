package Authentication

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	usersPath = `C:\Users\sokol\Desktop\Git\Library-on-go\Authentication\User`
	usersFile = "users.json"
)

func init() {
	if err := os.MkdirAll(usersPath, 0755); err != nil {
		fmt.Printf("Ошибка создания директории: %v\n", err)
	}
}

func (u *User) AddUser(name, gender, email, password string, age int) {
	u.Name = name
	u.Age = age
	u.Gender = gender
	u.Email = email
	u.Password = password
}

func getUsersFilePath() string {
	return filepath.Join(usersPath, usersFile)
}

func SaveAllUsers(users []User) error {
	filePath := getUsersFilePath()

	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка конвертации в JSON: %w", err)
	}

	// Сохраняем в файл
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка сохранения файла: %w", err)
	}

	fmt.Printf("Все пользователи сохранены в: %s\n", filePath)
	return nil
}

func LoadAllUsers() ([]User, error) {
	filePath := getUsersFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []User{}, nil // Возвращаем пустой список, если файла нет
	}

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var users []User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	return users, nil
}

func AddUserToFile(user User) error {
	if user.Name == "" && user.Email == "" && user.Password == "" {
		return fmt.Errorf("ОШИБКА: нельзя добавить пустого пользователя! Заполните данные пользователя")
	}

	users, err := LoadAllUsers()
	if err != nil {
		return err
	}

	filteredUsers := make([]User, 0, len(users)+1)
	for _, existingUser := range users {
		if existingUser.Name == "" && existingUser.Email == "" && existingUser.Password == "" {
			continue
		}
		filteredUsers = append(filteredUsers, existingUser)
	}

	filteredUsers = append(filteredUsers, user)

	return SaveAllUsers(filteredUsers)
}

func GetUsersPath() string {
	return usersPath
}

func EnsureDirectoryExists() error {
	return os.MkdirAll(usersPath, 0755)
}
