package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// Функция для шифрования пароля с использованием ключа
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Функция для проверки валидности пароля
func CheckPassword(hashedPassword, inputPassword string) bool {
	// Сравниваем хэш пароля и хэш введенного пароля
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
