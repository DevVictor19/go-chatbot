package controllers

import (
	"net/http"
	"server/src/api/dtos"
	"server/src/api/models"
	"server/src/api/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	businessRepository repositories.BusinessRepository
	customerRepository repositories.CustomerRepository
}

func (ctl *BusinessController) Create(ctx *gin.Context) {
	var body dtos.CreateBusinessDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	customer := ctl.getAuthenticatedCustomer(ctx)
	if customer == nil {
		return
	}

	now := time.Now()

	business := models.Business{
		CustomerId:  customer.ID,
		Name:        body.Name,
		Specialty:   body.Specialty,
		History:     body.History,
		ColorSchema: models.ColorSchema(body.ColorSchema),
		CreateAt:    now,
		UpdatedAt:   now,
	}

	if err := ctl.businessRepository.Create(ctx, &business); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (ctl *BusinessController) FindAllPaginated(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "0")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'page' parameter",
		})
		return
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'limit' parameter",
		})
		return
	}

	customer := ctl.getAuthenticatedCustomer(ctx)
	if customer == nil {
		return
	}

	results, err := ctl.businessRepository.FindAllPaginatedByCustomerId(
		ctx,
		customer.ID,
		page,
		limit,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": "Failed to find all paginated by customer id"})
		return
	}

	ctx.JSON(http.StatusOK, &results)
}

func (ctl *BusinessController) FindById(ctx *gin.Context) {
	businessId := ctx.Param("id")

	business, err := ctl.businessRepository.FindById(ctx, businessId)
	if err != nil || business == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Business not found"})
		return
	}

	customer := ctl.getAuthenticatedCustomer(ctx)
	if customer == nil {
		return
	}

	if business.CustomerId != customer.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Not allowed"})
		return
	}

	ctx.JSON(http.StatusOK, &business)
}

func (ctl *BusinessController) Update(ctx *gin.Context) {
	var body dtos.UpdateBusinessDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	businessId := ctx.Param("id")

	business, err := ctl.businessRepository.FindById(ctx, businessId)
	if err != nil || business == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Business not found"})
		return
	}

	customer := ctl.getAuthenticatedCustomer(ctx)
	if customer == nil {
		return
	}

	if business.CustomerId != customer.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Not allowed"})
		return
	}

	now := time.Now()

	business.ColorSchema = models.ColorSchema(body.ColorSchema)
	business.Name = body.Name
	business.History = body.History
	business.Specialty = body.Specialty
	business.UpdatedAt = now

	if err := ctl.businessRepository.Update(ctx, business); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}

func (ctl *BusinessController) Delete(ctx *gin.Context) {
	businessId := ctx.Param("id")

	business, err := ctl.businessRepository.FindById(ctx, businessId)
	if err != nil || business == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Business not found"})
		return
	}

	customer := ctl.getAuthenticatedCustomer(ctx)
	if customer == nil {
		return
	}

	if business.CustomerId != customer.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Not allowed"})
		return
	}

	if err := ctl.businessRepository.Delete(ctx, business.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}

func (ctl *BusinessController) getAuthenticatedCustomer(ctx *gin.Context) *models.Customer {
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

	customer, err := ctl.customerRepository.FindByEmail(ctx, emailString)
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

var businessController *BusinessController

func NewBusinessController(
	businessRepository *repositories.BusinessRepository,
	customerRepository *repositories.CustomerRepository) *BusinessController {

	if businessController == nil {
		businessController = &BusinessController{
			businessRepository: *businessRepository,
			customerRepository: *customerRepository,
		}
		return businessController
	}

	return businessController
}
