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

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		products := data.GetProducts()
		d, err := json.Marshal(products)
		if err != nil {
			http.Error(writer, "unable to marshal json", http.StatusInternalServerError)
		}
		writer.Write(d)
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)

}