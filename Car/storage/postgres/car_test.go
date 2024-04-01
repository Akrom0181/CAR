package postgres

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:  faker.Name(),
		Year:  2011,
		Brand: faker.Word(),
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	if assert.NoError(t, err) {
		createdCar, err := carRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, createdCar.Name)
			assert.Equal(t, reqCar.Year, createdCar.Year)
			assert.Equal(t, reqCar.Brand, createdCar.Brand)
		} else {
			return
		}
		fmt.Println("Created car", createdCar)
	}
}

func TestUpdateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Id:    "f9c1e1b8-7e67-4454-907a-c9d9c8138512",
		Name:  "M4 Competition",
		Brand: "BMW",
		Year:  2010,
	}

	id, err := carRepo.Update(context.Background(), reqCar)
	if assert.NoError(t, err) {
		updatedCar, err := carRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, updatedCar.Name)
			assert.Equal(t, reqCar.Brand, updatedCar.Brand)
			assert.Equal(t, reqCar.Year, updatedCar.Year)
		} else {
			return
		}
		fmt.Println("Updated car", updatedCar)
	}
}

func TestGetByIDCar(t *testing.T) {
	carRepo := NewCar(db)

	carID := "6c0ec5ed-6efe-482a-ba8b-24f3fce30e47"

	car, err := carRepo.GetByID(context.Background(), carID)

	if err != nil {
		t.Fatalf("error retrieving car with ID %s: %v", carID, err)
	}

	if car != (models.Car{}) {
		t.Errorf("expected nil car but got %+v when retrieving car with ID %s", car, carID)
	}
}

//// getAll ni chunmadim

func TestDeleteCar(t *testing.T) {
	carRepo := NewCar(db)

	carID := "d3a3ecf4-47bb-4c53-93ca-f63b7b72119c"

	err := carRepo.Delete(context.Background(), carID)

	if err == nil {
		t.Errorf("car with ID %s", carID)
	}

}
