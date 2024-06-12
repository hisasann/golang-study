// via https://engineering.mobalab.net/2021/10/04/where-should-interface-defined-in-go/

package main

import (
	"fmt"
	"log"

	"hisasann/golang-study/interface/user"
)

type mockUserRepository struct{}

func (r *mockUserRepository) CreateUser(user *user.User) error {
	return nil
}

type userCreator interface {
	CreateUser(user *user.User) error
}

// func RegisterUser(repo *user.Repository, email string) (*user.User, error) {
func RegisterUser(repo userCreator, email string) (*user.User, error) {
	u := &user.User{Email: email}
	if err := repo.CreateUser(u); err != nil {
		return nil, fmt.Errorf("unable to create user %w", err)
	}
	return u, nil
}

func main() {
	//creator := &mockUserRepository{}
	creator := &user.Repository{}
	_, err := RegisterUser(creator, "foo@example.com")
	log.Fatal(err)
}
