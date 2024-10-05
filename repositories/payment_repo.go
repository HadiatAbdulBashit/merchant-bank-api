package repositories

import (
	"merchant-bank-api/db"
)

func SavePayment(senderID, recipientID int, amount float64) error {
	query := "INSERT INTO payments (sender_id, recipient_id, amount) VALUES (?, ?, ?)"
	_, err := db.DB.Exec(query, senderID, recipientID, amount)
	return err
}
