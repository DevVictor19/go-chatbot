package controllers

import (
	"net/http"
	"server/src/api/dtos"
	"server/src/api/helpers"
	"server/src/api/models"
	"server/src/api/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository  *repositories.ProductRepository
	customerRepository *repositories.CustomerRepository
	businessRepository *repositories.BusinessRepository
}

func (ctl *ProductController) Create(ctx *gin.Context) {
	var body dtos.CreateProductDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	businessId := ctx.Param("businessId")

	helpers.CheckIfBusinessBelongsToCustomer(
		ctx,
		businessId,
		ctl.businessRepository,
		ctl.customerRepository,
	)

	now := time.Now()
	product := models.Product{
		BusinessId:    businessId,
		PhotoURL:      body.PhotoURL,
		Name:          body.Name,
		Description:   body.Description,
		StockQuantity: body.StockQuantity,
		Price:         body.Price,
		CreateAt:      now,
		UpdatedAt:     now,
	}

	if err := ctl.productRepository.Create(ctx, &product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (ctl *ProductController) FindAllPaginated(ctx *gin.Context) {
	businessId := ctx.Param("businessId")

	helpers.CheckIfBusinessBelongsToCustomer(
		ctx,
		businessId,
		ctl.businessRepository,
		ctl.customerRepository,
	)

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

	results, err := ctl.productRepository.FindAllPaginatedByBusinessId(
		ctx,
		businessId,
		page,
		limit,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": "Failed to find all paginated by business id"})
		return
	}

	ctx.JSON(http.StatusOK, &results)
}

func (ctl *ProductController) FindById(ctx *gin.Context) {
	businessId := ctx.Param("businessId")

	helpers.CheckIfBusinessBelongsToCustomer(
		ctx,
		businessId,
		ctl.businessRepository,
		ctl.customerRepository,
	)

	productId := ctx.Param("productId")

	product, err := ctl.productRepository.FindByIdAndBusinessId(
		ctx,
		productId,
		businessId,
	)
	if err != nil || product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, &product)
}

func (ctl *ProductController) Update(ctx *gin.Context) {
	var body dtos.UpdateProductDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	businessId := ctx.Param("businessId")

	helpers.CheckIfBusinessBelongsToCustomer(
		ctx,
		businessId,
		ctl.businessRepository,
		ctl.customerRepository,
	)

	productId := ctx.Param("productId")

	product, err := ctl.productRepository.FindByIdAndBusinessId(
		ctx,
		productId,
		businessId,
	)
	if err != nil || product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	now := time.Now()
	product.Description = body.Description
	product.Name = body.Name
	product.PhotoURL = body.PhotoURL
	product.Price = body.Price
	product.StockQuantity = body.StockQuantity
	product.UpdatedAt = now

	if err := ctl.productRepository.Update(ctx, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}

func (ctl *ProductController) Delete(ctx *gin.Context) {
	businessId := ctx.Param("businessId")

	helpers.CheckIfBusinessBelongsToCustomer(
		ctx,
		businessId,
		ctl.businessRepository,
		ctl.customerRepository,
	)

	productId := ctx.Param("productId")

	product, err := ctl.productRepository.FindByIdAndBusinessId(
		ctx,
		productId,
		businessId,
	)
	if err != nil || product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	if err := ctl.productRepository.DeleteByIdAndBusinessId(
		ctx,
		productId,
		businessId,
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}

var productController *ProductController

func NewProductController(
	productRepository *repositories.ProductRepository,
	customerRepository *repositories.CustomerRepository,
	businessRepository *repositories.BusinessRepository,
) *ProductController {
	if productController == nil {
		productController := &ProductController{
			productRepository:  productRepository,
			customerRepository: customerRepository,
			businessRepository: businessRepository,
		}
		return productController
	}

	return productController
}
