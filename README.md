# Merchant-Bank API

This is a backend API developed using Golang with the Gin framework, implementing essential functionalities for customer interactions between a merchant and a bank. The API includes login, registration, payment, and logout functionalities, with all activities logged in a MySQL database.

## Features

- **Registration**: Register a new customer.
- **Login**: Authenticate existing customers using email and password.
- **Payment**: Transfer payments between customers.
- **Logout**: Log out an authenticated customer.
- **Activity Logging**: All activities (login, payment, logout) are logged in the database.

## Technologies Used

- **Golang** (Gin framework) for the backend.
- **MySQL** for storing customer data and logs.
- **JWT (JSON Web Token)** for secure authentication.
- **Bcrypt** for password hashing.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/HadiatAbdulBashit/merchant-bank-api.git
   cd merchant-bank-api
   ```

2. **Create Database & Configure MySQL connection**

   run sql code to migrate database on db.sql

   Update the MySQL connection string in db/db.go file according to your MySQL configuration:

   ```go
   dsn := "<username>:<password>@tcp(127.0.0.1:<port>)/merchant_bank_api"
   ```

   example

   ```go
   dsn := "root:@tcp(127.0.0.1:3306)/merchant_bank_api"
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the project:**

   ```bash
   go run main.go
   ```

5. **Access the API:**

   The API will be running on localhost:8080. You can now make requests to the following endpoints:

   - POST /api/register: Register a new customer.
   - POST /api/login: Log in an existing customer and receive a JWT token.
   - POST /api/payment: Transfer a payment (requires JWT token).
   - POST /api/logout: Log out a customer (requires JWT token).
