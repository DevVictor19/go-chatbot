package controllers

import (
	"net/http"
	"server/src/api/dtos"
	"server/src/api/models"
	"server/src/api/repositories"
	"time"

	"server/src/api/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	repository repositories.CustomerRepository
}

func (c *AuthController) Login(ctx *gin.Context) {

}

func (c *AuthController) Signup(ctx *gin.Context) {
	var body dtos.SignupDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingCustomer, _ := c.repository.FindByEmail(ctx, body.Email)

	if existingCustomer != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Email already in use"})
		return
	}

	hashedPwd, err := services.Hash(body.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error while hashing password"})
		return
	}

	now := time.Now()

	customer := models.Customer{
		Name:      body.Name,
		Email:     body.Email,
		Password:  hashedPwd,
		CreateAt:  now,
		UpdatedAt: now,
	}

	c.repository.Create(ctx, &customer)
}

var authController *AuthController

func NewAuthController(customerRepository *repositories.CustomerRepository) *AuthController {
	if authController == nil {
		authController = &AuthController{
			repository: *customerRepository,
		}
		return authController
	}

	return authController
}
