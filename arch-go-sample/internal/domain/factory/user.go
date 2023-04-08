package factory

import "github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/model"

type User interface {
	// Create creates a user.
	Create(name string) (*model.User, error)
}
