// Package main runs the tavern and performs an Order
package main

import (
	"tavern/services/order"
	"tavern/services/tavern"

	"github.com/google/uuid"
)

func main() {

	products := productInventory()
	// Create Order Service to use in tavern
	os, err := order.NewOrderService(
		//order.WithMongoCustomerRepository("mongodb://admin:password@localhost:27017"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	// Create tavern service
	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os))
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
		products[0].GetID(),
		products[0].GetID(),
		products[0].GetID(),
		products[0].GetID(),
		products[0].GetID(),
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}
