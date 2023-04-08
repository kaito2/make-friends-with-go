package main

// go-chi/chi を利用してルーティングを行う
import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/handler"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/implement/factory"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/implement/repository"
	"github.com/kaito2/make-friends-with-go/arch-go-sample/internal/usecase"
)

func main() {
	// TODO: DI
	userRepository := repository.NewUserRepository()
	userFactory := factory.NewUser()
	registerUserUsecase := usecase.NewRegisterUser(userFactory, userRepository)
	registerUserHandler := handler.NewRegisterUserHandler(*registerUserUsecase)
	getUserHandler := handler.NewGetUserHandler(userRepository)

	// ルーティング
	r := chi.NewRouter()
	// http リクエストのログを出力
	r.Use(middleware.Logger)

	// ルーティング
	r.Route("/api/v1", func(r chi.Router) {
		// ユーザー
		r.Route("/users", func(r chi.Router) {
			r.Post("/", registerUserHandler.Handle)
			r.Get("/{user_id}", getUserHandler.Handle)
		})
	})

	// サーバー起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
