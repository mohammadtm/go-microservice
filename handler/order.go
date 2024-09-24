package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list order")
}
func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get order by id")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
