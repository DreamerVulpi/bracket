package main

import (
	"fmt"
	"log"
	"net/http"

	"context"

	"github.com/DreamerVulpi/bracket/config"
	"github.com/DreamerVulpi/bracket/handler"
	"github.com/DreamerVulpi/bracket/repo"
	"github.com/DreamerVulpi/bracket/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	cfg, err := config.LoadConfig("config/config.toml")
	if err != nil {
		log.Println(err)
		return
	}

	pool, err := pgxpool.New(context.Background(), cfg.Db.Dsn)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repo.User{Conn: pool}
	usecase := usecase.User{Repo: &repo}
	handler := &handler.Handler{UserUsecase: usecase}

	mux.HandleFunc("POST /api/v1/user", handler.AddUser)
	mux.HandleFunc("DELETE /api/v1/user", handler.DeleteUser)
	mux.HandleFunc("PATCH /api/v1/user", handler.EditUser)
	mux.HandleFunc("GET /api/v1/user", handler.GetUser)

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
