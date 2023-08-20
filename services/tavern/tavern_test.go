package tavern_test

import (
	"testing"

	"tavern/domain/product"
	"tavern/services/order"
	"tavern/services/tavern"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		//WithMemoryCustomerRepository(),
		order.WithMongoCustomerRepository("mongodb://admin:password@localhost:27017"),
		//WithMongoCustomerRepository("mongodb://admin:password@localhost:27017/?authMechanism=DEFAULT"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := tavern.NewTavern(tavern.WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	// cust, err := customer.NewCustomer("Percy")
	// if err != nil {
	// 	t.Error(err)
	// }

	// err = os.customers.Add(cust)
	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	//err = tavern.Order(cust.GetID(), order)
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}

}
