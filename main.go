package main

import (
	"fmt"
	"log"
	"net/http"

	"context"

	"github.com/DreamerVulpi/bracket/config"
	"github.com/DreamerVulpi/bracket/handler"
	"github.com/DreamerVulpi/bracket/pkg/jwt"
	"github.com/DreamerVulpi/bracket/repo"
	"github.com/DreamerVulpi/bracket/usecase"

	"github.com/DreamerVulpi/bracket/middleware"
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
	authUsecase := usecase.Auth{Repo: &repo.Auth{Conn: pool}, Jwt: jwt.Jwt{SecretKey: cfg.Jwt.Key}}
	middleware := middleware.Middleware{Jwt: jwt.Jwt{SecretKey: cfg.Jwt.Key}}
	handler := &handler.Handler{
		UserUsecase: userUsecase,
		SetUsecase:  setUsecase,
		PoolUsecase: poolUsecase,
		AuthUsecase: authUsecase,
	}

	r.HandleFunc("/api/v1/login", handler.Login)
	r.HandleFunc("/api/v1/register", handler.Register)

	r.HandleFunc("/api/v1/user", middleware.Auth(handler.AddUser))
	r.HandleFunc("/api/v1/user/{id}", middleware.Auth(handler.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/api/v1/user/{id}", middleware.Auth(handler.EditUser)).Methods("PATCH")
	r.HandleFunc("/api/v1/user/{id}", middleware.Auth(handler.GetUser)).Methods("GET")

	r.HandleFunc("/api/v1/set", middleware.Auth(handler.AddSet))
	r.HandleFunc("/api/v1/set/{id}", middleware.Auth(handler.DeleteSet)).Methods("DELETE")
	r.HandleFunc("/api/v1/set/{id}", middleware.Auth(handler.EditSet)).Methods("PATCH")
	r.HandleFunc("/api/v1/set/{id}", middleware.Auth(handler.GetSet)).Methods("GET")

	r.HandleFunc("/api/v1/pool", middleware.Auth(handler.AddPool))
	r.HandleFunc("/api/v1/pool/{id}", middleware.Auth(handler.DeletePool)).Methods("DELETE")
	r.HandleFunc("/api/v1/pool/{id}", middleware.Auth(handler.EditPool)).Methods("PATCH")
	r.HandleFunc("/api/v1/pool/{id}", middleware.Auth(handler.GetPool)).Methods("GET")

	fmt.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
