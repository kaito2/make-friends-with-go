package usecase

import (
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/factory"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/repository"
	"github.com/pkg/errors"
)

type (
	RegisterUser struct {
		userFactory    factory.User
		userRepository repository.User
	}
	RegisterUserInput struct {
		Name string
	}
)

// NewRegisterUser returns a new RegisterUser.
func NewRegisterUser(userFactory factory.User, userRepository repository.User) *RegisterUser {
	return &RegisterUser{
		userFactory:    userFactory,
		userRepository: userRepository,
	}
}

func (u *RegisterUser) Execute(input *RegisterUserInput) error {
	// ユーザーを作成する
	user, err := u.userFactory.Create(input.Name)
	if err != nil {
		return err
	}

	// ユーザーを保存する
	if err := u.userRepository.Put(user); err != nil {
		return errors.Wrap(err, "failed to put user")
	}

	return nil
}
