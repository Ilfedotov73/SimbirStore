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

type VendorService interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

type VendorProductService interface {
	GetProducts(ctx context.Context, page, size int, minPrice, maxPrice float64, vendorID int64, query, sort string) ([]model.Product, int, error)
}

type VendorHandler struct {
	vendorService  VendorService
	productService VendorProductService
}

func NewVendorHandler(vendorService VendorService, productService VendorProductService) *VendorHandler {
	return &VendorHandler{
		vendorService:  vendorService,
		productService: productService,
	}
}

func (h *VendorHandler) GetVendor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("vendorId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vendor id"})
		return
	}

	v, err := h.vendorService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Vendor not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vendorDTO := &dto.VendorDTO{
		ID:               v.ID,
		FirstName:        v.FirstName,
		LastName:         v.LastName,
		PhoneNumber:      v.PhoneNumber,
		PhotoURL:         v.PhotoURL,
		VendorTelegramID: v.UserTelegramID,
		CreateAt:         v.CreateAt,
	}
	c.JSON(http.StatusOK, vendorDTO)
}

func (h *VendorHandler) GetProductsByVendorID(c *gin.Context) {
	vendorID, err := strconv.ParseInt(c.Param("vendorId"), 10, 64)
	if err != nil || vendorID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vendor id"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	products, total, err := h.productService.GetProducts(c.Request.Context(), page, size, 0, 0, vendorID, "", "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dtos := make([]*dto.ProductSummary, len(products))
	for i := range products {
		d := toProductSummaryDTO(&products[i])
		dtos[i] = &d
	}

	c.JSON(http.StatusOK, dto.PagedProducts{
		Paging: dto.Paging{Page: page, Size: size, Total: total},
		Items:  dtos,
	})
}
