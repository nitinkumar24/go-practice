package handlers

import (
	"encoding/json"
	"go-practice/data"
	"log"
	"net/http"
)

type Products struct {
	l  *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	data, err := json.Marshal(products)
	if err != nil {
		http.Error(writer, "unable to marshal json", http.StatusInternalServerError)
	}
	writer.Write(data)
}