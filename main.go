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

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	r := mux.NewRouter()

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

	userUsecase := usecase.User{Repo: &repo.User{Conn: pool}}
	setUsecase := usecase.Set{Repo: &repo.Set{Conn: pool}}
	poolUsecase := usecase.Pool{Repo: &repo.Pool{Conn: pool}}
	authUsecase := usecase.Auth{Repo: &repo.Auth{Conn: pool}}

	handler := &handler.Handler{
		UserUsecase: userUsecase,
		SetUsecase:  setUsecase,
		PoolUsecase: poolUsecase,
		AuthUsecase: authUsecase,
		SecretKey:   cfg.Jwt.Key,
	}

	r.HandleFunc("/api/v1/login", handler.Login)
	r.HandleFunc("/api/v1/registration", handler.Register)

	r.HandleFunc("/api/v1/user", handler.AddUser)
	r.HandleFunc("/api/v1/user/{id}", handler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/v1/user/{id}", handler.EditUser).Methods("PATCH")
	r.HandleFunc("/api/v1/user/{id}", handler.GetUser).Methods("GET")

	r.HandleFunc("/api/v1/set", handler.AddSet)
	r.HandleFunc("/api/v1/set/{id}", handler.DeleteSet).Methods("DELETE")
	r.HandleFunc("/api/v1/set/{id}", handler.EditSet).Methods("PATCH")
	r.HandleFunc("/api/v1/set/{id}", handler.GetSet).Methods("GET")

	r.HandleFunc("/api/v1/pool", handler.AddPool)
	r.HandleFunc("/api/v1/pool/{id}", handler.DeletePool).Methods("DELETE")
	r.HandleFunc("/api/v1/pool/{id}", handler.EditPool).Methods("PATCH")
	r.HandleFunc("/api/v1/pool/{id}", handler.GetPool).Methods("GET")

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
