package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler отвечает "Hello, World!" на запросы к корневому маршруту
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
