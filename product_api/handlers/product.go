package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product_api/data"
	"strconv"
)

type Products struct {
	l  *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request)  {
	p.l.Println("handle post")
	prod := &data.Product{}
	err :=  prod.FromJSON(request.Body)
	if err !=  nil {
		http.Error(writer, "unable to  decode", http.StatusInternalServerError)
	}
	data.AddProduct(prod)
	p.l.Printf("body: %#v\n", prod)
}

func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request)  {
	products := data.GetProducts()
	d, err := json.Marshal(products)
	if err != nil {
		http.Error(writer, "unable to marshal json", http.StatusInternalServerError)
	}
	writer.Write(d)
}

func (p *Products) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "enable to parse id", http.StatusBadRequest)
	}
	p.l.Println("handle put", id)
	prod := &data.Product{}
	err =  prod.FromJSON(request.Body)
	if err !=  nil {
		http.Error(writer, "unable to  decode", http.StatusInternalServerError)
	}
	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(writer, "product not found", http.StatusNotFound)
	}
	p.l.Printf("body: %#v\n", prod)
}