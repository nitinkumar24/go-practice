package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID int					`json:"id"`
	Name string				`json:"name"`
	Description  string		`json:"description"`
	Price float32  			`json:"price"`
	SKU  string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID =  genNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	prod, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	// just for using returned values
	p.ID = id
	p.Name = "cycle"
	productList[pos] =  prod
	return err
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, fmt.Errorf("not found")
}

func genNextID()  int{
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

func GetProducts() [] *Product {
	return productList
}

var (
	productList = [] *Product{
		&Product{
			ID:          1,
			Name:        "Latte",
			Description: "Frothy milky coffee",
			Price:       2.45,
			SKU:         "abc323",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
		&Product{
			ID:          2,
			Name:        "Espresso",
			Description: "Short and strong coffee without milk",
			Price:       1.99,
			SKU:         "fjd34",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
	}
)