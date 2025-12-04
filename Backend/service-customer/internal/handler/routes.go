package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine, productHandler *ProductHandler, vendorHandler *VendorHandler, customerHandler *CustomerHandler, notificationHandler *NotificationHandler) {
	r.StaticFile("/openapi.yml", "./api/openapi.yml")

	docsUrl := ginSwagger.URL("/openapi.yml")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docsUrl))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	productGroup := r.Group("/products")
	{
		productGroup.GET("", productHandler.GetProducts)
		productGroup.GET("/:productId", productHandler.GetProductDetails)
		productGroup.GET("/:productId/reviews", productHandler.GetProductReviews)
	}

	vendorGroup := r.Group("/vendors")
	{
		vendorGroup.GET("/:vendorId", vendorHandler.GetVendor)
		vendorGroup.GET("/:vendorId/products", vendorHandler.GetProductsByVendorID)
	}

	customerGroup := r.Group("/customers")
	{
		customerGroup.GET("/:customerId", customerHandler.GetCustomer)
	}

	notificationGroup := r.Group("/notifications")
	{
		notificationGroup.GET("", notificationHandler.GetNotifications)
	}
}
