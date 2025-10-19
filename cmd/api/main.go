package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("GET /hello/{name}/{$}", helloHandler)

	if err := http.ListenAndServe(":8080", server); err != nil {
		slog.Error(slog.String("error", err.Error()).String())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	sparams := r.URL.Query()
	pathParam := r.PathValue("name")

	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")

	fmt.Println(databaseUser, databasePassword)

	text := sparams["teste"]
	w.Write([]byte(strings.Join(text, "/") + pathParam + databaseUser + databasePassword))
}
