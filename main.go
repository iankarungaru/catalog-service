package main

import(
	"fmt"
	"net/http"
)

func healthhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"status": "ok"}`)
}

func main() {
	http.HandleFunc("/health", healthhandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}