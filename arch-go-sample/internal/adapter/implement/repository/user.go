package repository

import (
	"fmt"
	"log"

	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/model"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/repository"
)

// NewUserRepository returns a user repository.
func NewUserRepository() repository.User {
	return &user{
		users: make(map[string]*model.User),
	}
}

// オンメモリでユーザーを管理する repository.User の実装
type user struct {
	users map[string]*model.User
}

// Get returns a user.
func (u *user) Get(id string) (*model.User, error) {
	user, ok := u.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// Put puts a user.
func (u *user) Put(user *model.User) error {
	log.Println("put user: ", user)

	u.users[user.ID.String()] = user

	return nil
}
