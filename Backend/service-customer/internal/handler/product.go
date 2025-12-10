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

type ProductService interface {
	GetFullProductData(ctx context.Context, productID int64) (*model.Product, *model.User, int, float64, error)
	GetProducts(ctx context.Context, page, size int, minPrice, maxPrice float64, vendorID int64, query, sort string) ([]model.Product, int, error)
}

type ReviewService interface {
	GetReviews(ctx context.Context, productID int64, page, size int) ([]model.Review, int, error)
}

type OfferService interface {
	CreateOffer(ctx context.Context, offer model.Offer) (*model.Offer, error)
}

type ProductHandler struct {
	productService ProductService
	reviewService  ReviewService
	offerService   OfferService
}

func NewProductHandler(productService ProductService, reviewService ReviewService, offerService OfferService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		reviewService:  reviewService,
		offerService:   offerService,
	}
}

func toProductSummaryDTO(p *model.Product) dto.ProductSummary {
	return dto.ProductSummary{
		ID:              p.ID,
		Name:            p.Name,
		Price:           p.Price,
		PhotoURL:        p.PhotoURL,
		Characteristics: p.Characteristics,
		ProductRating:   p.ProductRating,
		VendorID:        p.VendorID,
	}
}

func (h *ProductHandler) GetProductDetails(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("productId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	prod, vend, count, avg, err := h.productService.GetFullProductData(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prodDTO := &dto.ProductDetail{
		ProductSummary: toProductSummaryDTO(prod),
		CreateAt:       prod.CreateAt,
	}

	vendDTO := &dto.VendorDTO{
		ID:               vend.ID,
		FirstName:        vend.FirstName,
		LastName:         vend.LastName,
		PhoneNumber:      vend.PhoneNumber,
		PhotoURL:         vend.PhotoURL,
		VendorTelegramID: vend.UserTelegramID,
		CreateAt:         vend.CreateAt,
	}

	c.JSON(http.StatusOK, dto.FullProductResponse{
		Product:       prodDTO,
		Vendor:        vendDTO,
		ReviewsCount:  count,
		AverageRating: avg,
	})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("minPrice", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("maxPrice", "0"), 64)
	vendorID, _ := strconv.ParseInt(c.DefaultQuery("vendorId", "0"), 10, 64)
	query := c.Query("query")
	sort := c.Query("sort")

	products, total, err := h.productService.GetProducts(c.Request.Context(), page, size, minPrice, maxPrice, vendorID, query, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dtos := make([]*dto.ProductSummary, len(products))
	for i, p := range products {
		d := toProductSummaryDTO(&p)
		dtos[i] = &d
	}

	c.JSON(http.StatusOK, dto.PagedProducts{
		Paging: dto.Paging{Page: page, Size: size, Total: total},
		Items:  dtos,
	})
}

func (h *ProductHandler) GetProductReviews(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("productId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	items, total, err := h.reviewService.GetReviews(c.Request.Context(), id, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dtos := make([]*dto.ReviewDTO, len(items))
	for i, r := range items {
		dtos[i] = &dto.ReviewDTO{
			ID:         r.ID,
			CustomerID: r.CustomerID,
			ProductID:  r.ProductID,
			Review:     r.Review,
			Rating:     r.Rating,
		}
	}

	response := dto.PagedReviews{
		Paging: dto.Paging{Page: page, Size: size, Total: total},
		Items:  dtos,
	}
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CreateOffer(c *gin.Context) {
	productID, err := strconv.ParseInt(c.Param("productId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var req dto.OfferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offerModel := model.Offer{
		ProductID:  productID,
		VendorID:   req.VendorID,
		BuyerID:    req.BuyerID,
		OfferPrice: req.OfferPrice,
		Message:    req.Message,
		Status:     model.OfferStatusPending,
	}

	createdOffer, err := h.offerService.CreateOffer(c.Request.Context(), offerModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.OfferResponse{
		ID:         createdOffer.ID,
		ProductID:  createdOffer.ProductID,
		BuyerID:    createdOffer.BuyerID,
		VendorID:   createdOffer.VendorID,
		OfferPrice: createdOffer.OfferPrice,
		Message:    createdOffer.Message,
		Status:     string(createdOffer.Status),
		CreateAt:   createdOffer.CreateAt,
	}
	c.JSON(http.StatusCreated, resp)
}
