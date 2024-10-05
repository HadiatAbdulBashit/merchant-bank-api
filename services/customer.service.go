package services

import (
	"merchant-bank-api/models"
	"merchant-bank-api/repositories"
	"merchant-bank-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	customerID, err := utils.AuthenticateToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var paymentReq models.PaymentRequest
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	recipient, err := repositories.FindCustomerByID(paymentReq.RecipientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipient not found"})
		return
	}

	// Save payment
	err = repositories.SavePayment(customerID, recipient.ID, paymentReq.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	// Save payment history
	history := models.History{
		CustomerID: customerID,
		Action:     "payment",
		Timestamp:  time.Now(),
	}
	repositories.SaveHistory(history)

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
}
