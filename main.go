package main

import (
	"fmt"
	"net/http"
	

	"github.com/iankarungaru/catalog-service/handlers"
)





func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/products", handlers.ProductsHandler)
	mux.HandleFunc("/products/{id}", handlers.ProductDetailHandler)
	fmt.Println("Catalog service is running on port 8080...")
	http.ListenAndServe(":8080", handlers.LoggingMiddleware(mux))
}





func ApplyDiscount(price float64, discount float64) float64 {
	return price - (price * discount / 100)
}
