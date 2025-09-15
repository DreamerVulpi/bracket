package main

import (
	"fmt"
	"net/http"
	"os"

	"context"

	"github.com/DreamerVulpi/bracket/handler"
	"github.com/jackc/pgx/v5"
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
	mux.HandleFunc("POST /api/v1/user", handler.AddHandler)
	mux.HandleFunc("DELETE /api/v1/user", handler.DeleteHandler)
	mux.HandleFunc("PATCH /api/v1/user", handler.EditHandler)
	mux.HandleFunc("GET /api/v1/user", handler.GetHandler)

	// FIXME: Vulnarable
	conn, err := pgx.Connect(context.Background(), "postgresql://superuser:1234@localhost:5432/bracketProject")
	if err != nil {
		fmt.Printf("Unable to connect to database, %s", err)
		return
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Printf("QueryRow failed")
		return
	}

	fmt.Println(greeting)

	sqlBytes, err := os.ReadFile("init.sql")
	if err != nil {
		fmt.Printf("Read Sql file failed %s", err)
		return
	}

	_, err = conn.Exec(context.Background(), string(sqlBytes))
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
