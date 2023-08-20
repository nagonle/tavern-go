package order_test

import (
	"testing"

	"tavern/domain/product"
	"tavern/services/order"

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

func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// Add Customer
	//cust, err := customer.NewCustomer("Percy")
	//if err != nil {
	//t.Error(err)
	//}

	//err = os.customers.Add(cust)
	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	// Perform Order for one beer
	order := []uuid.UUID{
		products[0].GetID(),
	}

	//_, err = os.CreateOrder(cust.GetID(), order)
	_, err = os.CreateOrder(uid, order)

	if err != nil {
		t.Error(err)
	}

}
