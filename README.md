# congenial-octo-telegram

A simple users management system with concurrent

## Task

Необходимо разработать и реализовать web-сервис отвечающий за работу с пользователями. Главная особенность такого сервиса - данные пользователи хранятся в in-memory хранилище. Полученная реализация хранения пользователей должна учитывать, что запросы в сервис могут приходить параллельно и не должны ломать хранимые данные.

### Модель пользователя

- id - идентификатор, uuid
- name - ФИО
- login - логин пользователя
- password - пароль
- phone - телефон
- birthDate - дата рождения

### Сервис должен обрабатывать следующие запросы

- POST - /user - создание карточки пользователя
- PUT - /user - изменение карточки пользователя
- POST - /user/login - вход в систему. Данная ручка должна отдавать токен, позволяющий в дальнейшем идентифицировать пользователя, при работе с ручкой изменения, можно использовать внешние решения для работы с токенами, например jwt https://github.com/golang-jwt/jwt. Способ передачи токена - любой.

## How to use

- `go run cmd/main.go` - to run server
- ~~`go test -run . auth/server` - to run unit tests for server~~
- `go test -run . auth/store` - to run unit tests for store
- ~~`go test -run . auth/cmd` - to run interation tests~~
