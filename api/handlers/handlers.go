package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"

	"udacityGoCRMBackend/api/database"
	"udacityGoCRMBackend/api/database/models"
)

// TODO: Add Docstrings to all functions/helpers and optional comments to explain code
// TODO: Refactor handlers: split up and maybe add helper functions
// TODO: Finish README file content
// TODO: Rename Project
// TODO: Add project to git

// TODO: Add new endpoint for batch updating
// TODO: Replace mock database with real Database

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/web/index.html")
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	customerFound := false
	var matchedCustomer models.Customer

	for _, customer := range database.CustomerDb {
		if strconv.Itoa(customer.Id) == vars["id"] {
			customerFound = true
			matchedCustomer = customer
		}
	}

	if customerFound {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Getting customer data with ID", matchedCustomer.Id)
		err := json.NewEncoder(w).Encode(matchedCustomer)
		handleError(err)
	} else {
		customerNotExists(w)
	}
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	fmt.Println("Getting all customer data")
	err := json.NewEncoder(w).Encode(database.CustomerDb)
	handleError(err)
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer models.Customer
	customerIdExists := false

	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &newCustomer)
	handleError(err)

	for _, customer := range database.CustomerDb {
		if customer.Id == newCustomer.Id {
			customerIdExists = true
		}
	}

	if customerIdExists {
		w.WriteHeader(http.StatusConflict)
		_, err = fmt.Fprintf(w, "User with this ID already exists..\n")
		handleError(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Println("Created new customer with ID", newCustomer.Id)
		database.CustomerDb = append(database.CustomerDb, newCustomer)

		err = json.NewEncoder(w).Encode(database.CustomerDb)
		handleError(err)
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	customerFound := false
	var customerIndex int
	var updatedCustomer models.Customer

	for index, customer := range database.CustomerDb {
		if strconv.Itoa(customer.Id) == vars["id"] {
			customerFound = true
			customerIndex = index
		}
	}

	if customerFound {
		reqBody, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(reqBody, &updatedCustomer)
		handleError(err)

		//TODO: Refactor update logic
		var existingCustomer *models.Customer = &database.CustomerDb[customerIndex]
		existingCustomer.Name = updatedCustomer.Name
		existingCustomer.Role = updatedCustomer.Role
		existingCustomer.Email = updatedCustomer.Email
		existingCustomer.Phone = updatedCustomer.Phone
		existingCustomer.Contacted = updatedCustomer.Contacted

		w.WriteHeader(http.StatusCreated)
		fmt.Println("Updated customer with ID", existingCustomer.Id)

		err = json.NewEncoder(w).Encode(database.CustomerDb)
		handleError(err)
	} else {
		customerNotExists(w)
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	customerFound := false
	var customerIndex int

	for index, customer := range database.CustomerDb {
		if strconv.Itoa(customer.Id) == vars["id"] {
			customerFound = true
			customerIndex = index
		}
	}

	if customerFound {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Deleted customer with ID", database.CustomerDb[customerIndex].Id)
		database.CustomerDb = append(database.CustomerDb[:customerIndex], database.CustomerDb[customerIndex+1:]...)

		err := json.NewEncoder(w).Encode(database.CustomerDb)
		handleError(err)
	} else {
		customerNotExists(w)
	}
}
