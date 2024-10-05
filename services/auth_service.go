package services

import (
	"merchant-bank-api/models"
	"merchant-bank-api/repositories"
	"merchant-bank-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerReq models.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the customer already exists
	_, err := repositories.FindCustomerByEmail(registerReq.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer already exists"})
		return
	}

	// Create the customer
	err = repositories.CreateCustomer(registerReq.Email, registerReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer registered successfully"})
}

func Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	customer, err := repositories.FindCustomerByEmail(loginReq.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(loginReq.Password, customer.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Save login history
	history := models.History{
		CustomerID: customer.ID,
		Action:     "login",
		Timestamp:  time.Now(),
	}
	repositories.SaveHistory(history)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func Logout(c *gin.Context) {
	customerID, err := utils.AuthenticateToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Save logout history
	history := models.History{
		CustomerID: customerID,
		Action:     "logout",
		Timestamp:  time.Now(),
	}
	repositories.SaveHistory(history)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
