package postgres

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"testing"


	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	orderRepo := NewOrder(db)

	reqOrder := models.CreateOrderr{
		FromDate: "2024-02-18",
	}

	id, err := orderRepo.CreateOrder(context.Background(), reqOrder)
	if assert.NoError(t, err) {
		createdOrder, err := orderRepo.GetOne(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqOrder.FromDate, createdOrder.FromDate)
		} else {
			return
		}
		fmt.Println("Created customer", createdOrder)
	}
}

func TestUpdateOrder(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.Customer{
		FirstName: "Imron",
	}
	updateAt, err := customerRepo.Update(context.Background(), reqCustomer)
	if assert.NoError(t, err) {
		updatedCustomer, err := customerRepo.GetByID(context.Background(), updateAt)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCustomer.UpdatedAt, updatedCustomer.UpdatedAt)
		} else {
			return
		}
		fmt.Println("updated customer", updatedCustomer)
	}
}

func TestGetByIDOrder(t *testing.T) {
	customerRepo := NewOrder(db)

	orderID := "bf767d11-5cc8-43cc-85d1-54ed587fc32e"

	car, err := customerRepo.GetOne(context.Background(), orderID)

	if err != nil {
		t.Fatalf("error retrieving customer with ID %s: %v", orderID, err)
	}

	if car != (models.GetOrder{}) {
		t.Errorf("expected nil car but got %+v when retrieving customer with ID %s", car, orderID)
	}
}


func TestDeleteOrder(t *testing.T) {
	orderRepo := NewOrder(db)

	orderID := "bf767d11-5cc8-43cc-85d1-54ed587fc32e"

	err := orderRepo.DeleteOrder(context.Background(), orderID)

	if err == nil {
		t.Errorf("order with ID %s", orderID)
	}

}