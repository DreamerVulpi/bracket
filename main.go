package main

import (
	"fmt"
	"net/http"
	"os"

	"context"

	"github.com/DreamerVulpi/bracket/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

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

	// FIXME: Vulnarable
	pool, err := pgxpool.New(context.Background(), "postgresql://superuser:1234@localhost:5432/bracketProject")

	handler := &handler.Handler{}

	mux.HandleFunc("POST /api/v1/user", handler.AddUser)
	mux.HandleFunc("DELETE /api/v1/user", handler.DeleteUser)
	mux.HandleFunc("PATCH /api/v1/user", handler.EditUser)
	mux.HandleFunc("GET /api/v1/user", handler.GetUser)

	sqlBytes, err := os.ReadFile("init.sql")
	if err != nil {
		fmt.Printf("Read Sql file failed %s", err)
		return
	}

	_, err = pool.Exec(context.Background(), string(sqlBytes))
	if err != nil {
		fmt.Printf("Can't init database %s", err)
	}

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
