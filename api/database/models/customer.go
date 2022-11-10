package models

// Customer is a model for customers in the mock database.
type Customer struct {
	Id        int
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}
