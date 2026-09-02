package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(products)
		return
	}
	if r.Method == http.MethodPost {
		createProductHandler(w, r)
		return

	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/health", healthhandler)
	http.HandleFunc("/products", productsHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
		return
	}
	products = append(products, newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}
