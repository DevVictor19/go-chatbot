package helpers

import (
	"net/http"
	"server/src/api/repositories"

	"github.com/gin-gonic/gin"
)

func CheckIfBusinessBelongsToCustomer(
	ctx *gin.Context,
	businessId string,
	businessRepository *repositories.BusinessRepository,
	customerRepository *repositories.CustomerRepository,
) {
	business, err := businessRepository.FindById(ctx, businessId)
	if err != nil || business == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Business not found"})
		return
	}

	customer := GetAuthenticatedCustomer(ctx, customerRepository)
	if customer == nil {
		return
	}

	if business.CustomerId != customer.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Not allowed"})
		return
	}
}
