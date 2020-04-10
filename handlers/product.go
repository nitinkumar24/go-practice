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

func (p *Products) addProduct(writer http.ResponseWriter, request *http.Request)  {
	p.l.Println("handle post")
	prod := &data.Product{}
	err :=  prod.FromJSON(request.Body)
	if err !=  nil {
		http.Error(writer, "unable to  decode", http.StatusInternalServerError)
	}
	data.AddProduct(prod)
	p.l.Printf("body: %#v\n", prod)
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request)  {
	products := data.GetProducts()
	d, err := json.Marshal(products)
	if err != nil {
		http.Error(writer, "unable to marshal json", http.StatusInternalServerError)
	}
	writer.Write(d)
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(writer, request)
		return
	}

	if request.Method == http.MethodPost {
		p.addProduct(writer,  request)
		return
	}


	writer.WriteHeader(http.StatusMethodNotAllowed)
}