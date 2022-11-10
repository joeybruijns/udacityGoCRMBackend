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

// Home serves a static html home page.
func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/web/index.html")
}

// GetCustomer searches the database for a user with a given ID in the HTTP GET request.
// It returns a 200 OK status if the customer was found.
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
		handleBasicError(err)
	} else {
		customerNotExists(w)
	}
}

// GetCustomers returns all the customers from the database.
// It gives a 200 OK status if the request was made successfully.
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	fmt.Println("Getting all customer data")

	err := json.NewEncoder(w).Encode(database.CustomerDb)
	handleBasicError(err)
}

// AddCustomer handles an HTTP POST request for creating a new customer in the database.
// It gives back a 201 Created status if the customer is created successfully,
// and a 409 Conflict status if a user with the provided ID already exists.
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer models.Customer
	customerFound := false

	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &newCustomer)
	handleBasicError(err)

	for _, customer := range database.CustomerDb {
		if customer.Id == newCustomer.Id {
			customerFound = true
		}
	}

	if customerFound {
		w.WriteHeader(http.StatusConflict)
		_, err = fmt.Fprintf(w, "User with this ID already exists..\n")
		handleBasicError(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Println("Created new customer with ID", newCustomer.Id)
		database.CustomerDb = append(database.CustomerDb, newCustomer)

		err = json.NewEncoder(w).Encode(database.CustomerDb)
		handleBasicError(err)
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
		handleBasicError(err)

		//Overwrite data of existing customer except for the ID value
		var existingCustomer *models.Customer = &database.CustomerDb[customerIndex]
		existingCustomer.Name = updatedCustomer.Name
		existingCustomer.Role = updatedCustomer.Role
		existingCustomer.Email = updatedCustomer.Email
		existingCustomer.Phone = updatedCustomer.Phone
		existingCustomer.Contacted = updatedCustomer.Contacted

		w.WriteHeader(http.StatusCreated)
		fmt.Println("Updated customer with ID", existingCustomer.Id)

		err = json.NewEncoder(w).Encode(database.CustomerDb)
		handleBasicError(err)
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
		handleBasicError(err)
	} else {
		customerNotExists(w)
	}
}
