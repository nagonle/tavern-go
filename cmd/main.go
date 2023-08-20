package main

import (
	"fmt"
	"tavern/services/tavern"

	"tavern/domain/product"

	//"tavern/domain/product"
	"tavern/services/order"

	"github.com/google/uuid"
)

func main() {
	products := productInventory()

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

	//cust, err := customer.NewCustomer("Percy")
	//cust, err := os.NewCustomer("Percy")
	cust, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}

	//uid := cust.GetID()
	// uid_str := "64de0c44048b401a11545912"
	// uid, _ := uuid.Parse(uid_str)
	//fmt.Println("uid:", uid)
	fmt.Println("uid:", cust)

	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}
