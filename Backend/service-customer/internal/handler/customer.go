package handler

import (
	"context"
	"database/sql"
	"net/http"
	"service-customer/internal/dto"
	"service-customer/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerService interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

type CustomerHandler struct {
	customerService CustomerService
}

func NewCustomerHandler(customerService CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("customerId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer id"})
		return
	}

	customer, err := h.customerService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customerDTO := &dto.CustomerDTO{
		ID:                 customer.ID,
		FirstName:          customer.FirstName,
		LastName:           customer.LastName,
		PhoneNumber:        customer.PhoneNumber,
		PhotoURL:           customer.PhotoURL,
		CustomerTelegramID: customer.UserTelegramID,
		CreateAt:           customer.CreateAt,
		Login:              customer.Login,
		Email:              customer.Email,
	}
	c.JSON(http.StatusOK, customerDTO)
}
