# Udacity - Go CRM Backend Project

This application is the backend of a CRM system. 
Users will be able to make HTTP requests to the server to perform basic CRUD operations. 
A user can get all the customers from the mock database, get a single user by ID, create a new user, 
update an existing user or delete a user by ID.  

### Installing and Launching the Project
To launch the project, run the following command from the root directory of the project:
`go run main.go`. The application runs on localhost port 3000.

### Basic Usage
This is an overview of the available operations that are supported, together with the available endpoints:
- (**GET**) Get a single customer: `/customers/{id}`
- (**GET**) Get all customers: `/customers`
- (**POST**) Create a new customer: `/customers`
- (**PUT**) Update an existing customer: `/customers/{id}`
- (**DELETE**) Delete a customer: `/customers/{id}`

**Example:** curl command - POST request for creating a new user:
```shell
curl --location --request POST 'http://localhost:3000/customers' \
--header 'Content-Type: application/json' \
--data-raw '    {
        "Id": 4,
        "Name": "Jon",
        "Role": "Customer",
        "Email": "jon@jon.com",
        "Phone": "1234543210",
        "Contacted": false
    }'
```