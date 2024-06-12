package user

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID    int
	Email string
}

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateUser(user *User) error {
	fmt.Println(user.Email, "created")
	return nil
}
