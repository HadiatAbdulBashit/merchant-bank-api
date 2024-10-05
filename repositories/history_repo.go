package repositories

import (
	"merchant-bank-api/db"
	"merchant-bank-api/models"
)

func SaveHistory(entry models.History) error {
	query := "INSERT INTO history (customer_id, action, timestamp) VALUES (?, ?, ?)"
	_, err := db.DB.Exec(query, entry.CustomerID, entry.Action, entry.Timestamp)
	return err
}
