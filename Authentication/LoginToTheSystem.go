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
	key := GetKey(16)
	u.AuthenticationCode = key
	fmt.Println("Ваш код аундефикации для входа в библиотеку: ", u.AuthenticationCode)
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

func SimilarityCheck(user *User) error {
	filePath := getUsersFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Файл с пользователями не найден, email уникален")
		return nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if len(data) == 0 {
		fmt.Println("Файл пустой, email уникален")
		return nil
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	for _, existingUser := range users {
		if existingUser.Email == user.Email {
			return fmt.Errorf("ОШИБКА: пользователь с email '%s' уже существует", user.Email)
		}
	}
	return nil
}

func AddUserToFile(user User) error {
	// Проверяем, что пользователь не пустой
	if user.Name == "" && user.Email == "" && user.Password == "" {
		return fmt.Errorf("ОШИБКА: нельзя добавить пустого пользователя! Заполните данные пользователя")
	}

	// Проверяем уникальность email (ВАЖНО: обрабатываем ошибку!)
	err := SimilarityCheck(&user)
	if err != nil {
		return err // Возвращаем ошибку, если email уже существует
	}

	// Загружаем существующих пользователей
	users, err := LoadAllUsers()
	if err != nil {
		return err
	}

	// Фильтруем пустых пользователей
	filteredUsers := make([]User, 0, len(users)+1)
	for _, existingUser := range users {
		if existingUser.Name == "" && existingUser.Email == "" && existingUser.Password == "" {
			continue
		}
		filteredUsers = append(filteredUsers, existingUser)
	}

	// Добавляем нового пользователя
	filteredUsers = append(filteredUsers, user)

	// Сохраняем всех пользователей
	return SaveAllUsers(filteredUsers)
}

func GetUsersPath() string {
	return usersPath
}

func EnsureDirectoryExists() error {
	return os.MkdirAll(usersPath, 0755)
}
