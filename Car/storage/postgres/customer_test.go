package postgres

import (
	"context"
	"fmt"
	"rent-customer/api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.Customer{
		FirstName:  "Imron",
		LastName:   "Hojiyev",
		Gmail:      "imronhojiyev@gmail.com",
		Phone:      "+998999827640",
		Is_Blocked: false,
	}

	id, err := customerRepo.Create(context.Background(), reqCustomer)
	if assert.NoError(t, err) {
		updatedCustomer, err := customerRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCustomer.FirstName, updatedCustomer.FirstName)
			assert.Equal(t, reqCustomer.LastName, updatedCustomer.LastName)
			assert.Equal(t, reqCustomer.Gmail, updatedCustomer.Gmail)
			assert.Equal(t, reqCustomer.Phone, updatedCustomer.Phone)
			assert.Equal(t, reqCustomer.Is_Blocked, updatedCustomer.Is_Blocked)

		} else {
			return
		}
		fmt.Println("Created customer", updatedCustomer)
	}
}

func TestUpdateCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.Customer{
		Id: "07d7b1a7-3945-43b0-b722-f96fdf47098a",
		FirstName:  "Imron11",
		LastName:   "Timurov1111",
		Gmail:      "imrontimurov@gmail.com",
		Phone:      "+998999827649",
		Is_Blocked: false,
	}

	id, err := customerRepo.Update(context.Background(), reqCustomer)
	if assert.NoError(t, err) {
		updatedCustomer, err := customerRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCustomer.FirstName, updatedCustomer.FirstName)
			assert.Equal(t, reqCustomer.LastName, updatedCustomer.LastName)
			assert.Equal(t, reqCustomer.Gmail, updatedCustomer.Gmail)
			assert.Equal(t, reqCustomer.Phone, updatedCustomer.Phone)
			assert.Equal(t, reqCustomer.Is_Blocked, updatedCustomer.Is_Blocked)

		} else {
			return
		}
		fmt.Println("Updated customer", updatedCustomer)
	}
}

func TestGetByIDCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	customerID := "a6824eb8-0d5c-4072-a79f-34f09062c717"

	customer, err := customerRepo.GetByID(context.Background(), customerID)

	if err != nil {
		t.Fatalf("error retrieving customer with ID %s: %v", customerID, err)
	}

	if customer != (models.Customer{}) {
		t.Errorf("expected nil customer but got %+v when retrieving customer with ID %s", customer, customerID)
	}
}


func TestDeleteCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	customerID := "57d613c7-7a7f-45f6-976e-1c6612cd4b10"

	err := customerRepo.Delete(context.Background(), customerID)

	if err == nil {
		t.Errorf("customer with ID %s", customerID)
	}

}