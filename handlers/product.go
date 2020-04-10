package handlers

import (
	"encoding/json"
	"go-practice/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if request.Method == http.MethodPut {
		p.l.Println("method put", request.URL.Path)
		path := request.URL.Path
		r := regexp.MustCompile(`/([0-9]+)`)
		group := r.FindAllStringSubmatch(path, -1)
		if len(group) !=1 {
			http.Error(writer, "not found", http.StatusBadRequest)
		}
		if len(group[0]) !=1 {
			http.Error(writer, "not found", http.StatusBadRequest)
		}

		idString  := group[0][1]
		id, _ := strconv.Atoi(idString)
		p.l.Println("id",  id)
		p.updateProduct(writer, request, id)
		return
	}
	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProduct(writer http.ResponseWriter, request *http.Request, id int) {
	p.l.Println("handle put")
	prod := &data.Product{}
	err :=  prod.FromJSON(request.Body)
	if err !=  nil {
		http.Error(writer, "unable to  decode", http.StatusInternalServerError)
	}
	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(writer, "product not found", http.StatusNotFound)
	}
	p.l.Printf("body: %#v\n", prod)
}