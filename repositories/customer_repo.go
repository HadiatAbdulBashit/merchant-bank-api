package repositories

import (
	"database/sql"
	"errors"
	"merchant-bank-api/db"
	"merchant-bank-api/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateCustomer(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "INSERT INTO customers (email, password) VALUES (?, ?)"
	_, err = db.DB.Exec(query, email, hashedPassword)
	return err
}

func FindCustomerByEmail(email string) (models.Customer, error) {
	var customer models.Customer
	query := "SELECT id, email, password FROM customers WHERE email = ?"
	err := db.DB.QueryRow(query, email).Scan(&customer.ID, &customer.Email, &customer.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, errors.New("Customer not found")
		}
		return customer, err
	}
	return customer, nil
}

func FindCustomerByID(id int) (models.Customer, error) {
	var customer models.Customer
	query := "SELECT id, email, password FROM customers WHERE id = ?"
	err := db.DB.QueryRow(query, id).Scan(&customer.ID, &customer.Email, &customer.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, errors.New("Customer not found")
		}
		return customer, err
	}
	return customer, nil
}
