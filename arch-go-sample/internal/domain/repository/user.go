package repository

import "github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/model"

type User interface {
	// Get returns a user.
	Get(id string) (*model.User, error)
	// Put puts a user.
	Put(user *model.User) error
}
