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
	id        uuid.UUID
	name      string
	login     string
	password  string
	phone     string
	birthDate time.Time
}

func NewUser(name, login, password, phone string, date time.Time) *User {
	return &User{
		id:        uuid.New(),
		name:      name,
		login:     login,
		password:  password,
		phone:     phone,
		birthDate: date,
	}
}
