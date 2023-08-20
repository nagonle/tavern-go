package services

import (
	"testing"

	"tavern/domain/customer"

	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := NewOrderService(
		//WithMemoryCustomerRepository(),
		WithMongoCustomerRepository("mongodb://admin:password@localhost:27017"),
		//WithMongoCustomerRepository("mongodb://admin:password@localhost:27017/?authMechanism=DEFAULT"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
