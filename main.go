package main

import (
	"fmt"
	"net/http"

	"github.com/DreamerVulpi/bracket/handler"
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

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
