package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"udacityGoCRMBackend/api/database"
	"udacityGoCRMBackend/api/handlers"
)

func main() {
	database.SeedCustomerData()

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/customers/{id}", handlers.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers", handlers.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000..")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
