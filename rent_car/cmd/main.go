package main

import (
	"fmt"
	"rent_car/config"
	"rent_car/controller"
	"rent_car/models"
	"rent_car/storage"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	c := controller.NewController(store)
	// c.CreateCar()
	// c.UpdateCar()
	c.Store.Car.Delete(models.Car{},"845483f3-787c-4d06-bf33-e22277765edc")
	// c.Store.Car.GetAll()

}
