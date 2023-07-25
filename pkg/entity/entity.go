package entity

import (
	"encoding/json"
	"log"
)

type ProductInterface interface {
	String() string
}

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type ProductList struct {
	List []*Product `json:"list"`
}

func (p *Product) String() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err.Error())
	}

	return string(data)
}

func NewProduct(name, code string, price float64) *Product {
	return &Product{
		Name:  name,
		Code:  code,
		Price: price,
	}
}
