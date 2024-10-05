package models

type PaymentRequest struct {
	RecipientID int     `json:"recipient_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
}
