package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Product is the struct
type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJSON := `[
		{"productId": 3, "manufacturer": "Sas", "sku": "Asdad", "upc": "12313131", "pricePerUnit": "131.12", "quantityOnHand": 5908, "productName": "la"}
	]`

	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)
	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
