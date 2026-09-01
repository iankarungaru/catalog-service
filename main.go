package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/iankarungaru/catalog-service/models"
)
var products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 999.99, Stock: 10},
	{ID: 2, Name: "Smartphone", Price: 499.99, Stock: 20},
	{ID: 3, Name: "Tablet", Price: 299.99, Stock: 15},
}

func healthhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"status": "ok"}`)
}
func productsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/health", healthhandler)
	http.HandleFunc("/products", productsHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}