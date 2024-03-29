package main

import (
	"context"
	"fmt"
	"rent-car/api"
	"rent-car/config"
	"rent-car/service"
	"rent-car/storage/postgres"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	services := service.New(store)

	c := api.New(services, store)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8008")
}
