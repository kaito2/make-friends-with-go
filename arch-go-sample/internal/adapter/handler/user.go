package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/repository"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/usecase"
)

// RegisterUser usecase.RegisterUser を実行する net/http.Handler
type RegisterUserHandler struct {
	usecase usecase.RegisterUser
}

// NewRegisterUser RegisterUser を生成する
func NewRegisterUserHandler(usecase usecase.RegisterUser) *RegisterUserHandler {
	return &RegisterUserHandler{usecase: usecase}
}

// Handle RegisterUser を実行する
func (h *RegisterUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if err := h.usecase.Execute(&usecase.RegisterUserInput{Name: name}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type GetUserHandler struct {
	// TODO: repository を使わないようにする
	userRepository repository.User
}

func NewGetUserHandler(userRepository repository.User) *GetUserHandler {
	return &GetUserHandler{userRepository: userRepository}
}

func (h *GetUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// id に該当するユーザーを h.userRepository を使って取得する
	id := chi.URLParam(r, "user_id")
	user, err := h.userRepository.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
