package models

import (
	"time"

	"github.com/google/uuid"
)

/*
- id - идентификатор, uuid
- name - ФИО
- login - логин пользователя
- password - пароль
- phone - телефон
- birthDate - дата рождения
*/

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birthDate"`
}

func NewUser(name, login, password, phone string, date time.Time) *User {
	return &User{
		ID:        uuid.New(),
		Name:      name,
		Login:     login,
		Password:  password,
		Phone:     phone,
		BirthDate: date,
	}
}

func (u *User) Patch(other *User) {
	u.Name = other.Name
	u.Login = other.Login
	u.Password = other.Password
	u.Phone = other.Phone
	u.BirthDate = other.BirthDate
}
