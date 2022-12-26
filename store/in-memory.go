package store

type UsersStore interface {
	CreateUserRecord(name string)
	UpdateUserRecord(name, patch string)
}
