package helpers

import (
	"net/http"
	"server/src/api/models"
	"server/src/api/repositories"

	"github.com/gin-gonic/gin"
)

func GetAuthenticatedCustomer(
	ctx *gin.Context,
	customerRepository *repositories.CustomerRepository) *models.Customer {

	email, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Email not found in context"})
		return nil
	}

	emailString, ok := email.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid type assertion email"})
		return nil
	}

	customer, err := customerRepository.FindByEmail(ctx, emailString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to find customer by email"})
		return nil
	}

	if customer == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Customer not found"})
		return nil
	}

	return customer
}
