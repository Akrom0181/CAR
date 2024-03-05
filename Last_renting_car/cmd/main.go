package main

import (
	"fmt"
	"net/http"
	"rent-car/config"
	"rent-car/controller"
	"rent-car/storage"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	con := controller.NewController(store)

	http.HandleFunc("/car", con.Car)
	http.HandleFunc("/customer", con.Customer)

	fmt.Println("programm is running on localhost:9090...")
	http.ListenAndServe(":9090", nil)
	
}
