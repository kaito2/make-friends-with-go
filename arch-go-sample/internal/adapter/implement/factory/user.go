package factory

import (
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/factory"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/model"
	"github.com/rs/xid"
)

type user struct{}

// NewUser returns a new User.
func NewUser() factory.User {
	return &user{}
}

func (u *user) Create(name string) (*model.User, error) {
	// ユーザーを作成する
	user := &model.User{
		ID:   xid.New(),
		Name: name,
	}

	return user, nil
}
