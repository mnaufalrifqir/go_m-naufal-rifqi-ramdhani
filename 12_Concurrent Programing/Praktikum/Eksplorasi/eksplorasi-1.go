package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Products struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func main() {
	var products []Products
	url := "https://fakestoreapi.com/products"
	ch := make(chan []Products)

	go func() {
		resp, err := http.Get(url)
		if err == nil {
			err1 := json.NewDecoder(resp.Body).Decode(&products)
			if err1 == nil {
				ch <- products
			} else {
				fmt.Println("Error:", err1)
				ch <- nil
			}
		} else {
			fmt.Println("Error:", err)
			ch <- nil

		}
		resp.Body.Close()
	}()
	products = <-ch
	if products != nil {
		fmt.Println("Products Data")
		for _, product := range products {
			fmt.Printf("===\nTitle: %s\nPrice: %.2f\nCategory: %s\n===\n", product.Title, product.Price, product.Category)
		}
	} else {
		fmt.Println("List Product Kosong")
	}
}
