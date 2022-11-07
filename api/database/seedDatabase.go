package database

import "udacityGoCRMBackend/api/database/models"

func SeedCustomerData() {
	c1 := models.Customer{
		Id:        1,
		Name:      "Bob",
		Role:      "Customer",
		Email:     "bob@bob.com",
		Phone:     "1234567890",
		Contacted: false,
	}

	c2 := models.Customer{
		Id:        2,
		Name:      "Jane",
		Role:      "Customer",
		Email:     "jane@jane.com",
		Phone:     "9876543210",
		Contacted: true,
	}

	c3 := models.Customer{
		Id:        3,
		Name:      "May",
		Role:      "Customer",
		Email:     "may@may.com",
		Phone:     "4753890912",
		Contacted: true,
	}

	CustomerDb = []models.Customer{
		c1,
		c2,
		c3,
	}
}
