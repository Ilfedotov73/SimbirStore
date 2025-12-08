package app

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"service-customer/internal/handler"
	"service-customer/internal/middleware"
	"service-customer/internal/repository"
	"service-customer/internal/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/service_customer_db?sslmode=disable"
	}

	db, err := newDB(dsn)
	if err != nil {
		logger.Error("failed to init db", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	productRepository := repository.NewProductRepository(db)
	vendorRepository := repository.NewVendorRepository(db)
	reviewRepository := repository.NewReviewRepository(db)
	customerRepository := repository.NewCustomerRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	offerRepository := repository.NewOfferRepository(db)

	vendorService := service.NewVendorService(vendorRepository)
	reviewService := service.NewReviewService(reviewRepository)
	productService := service.NewProductService(productRepository, vendorService, reviewService)
	customerService := service.NewCustomerService(customerRepository)
	notificationService := service.NewNotificationService(notificationRepository)
	offerService := service.NewOfferService(offerRepository)

	productHandler := handler.NewProductHandler(productService, reviewService, offerService)
	vendorHandler := handler.NewVendorHandler(vendorService, productService)
	customerHandler := handler.NewCustomerHandler(customerService)
	notificationHandler := handler.NewNotificationHandler(notificationService)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	handler.InitRoutes(router, productHandler, vendorHandler, customerHandler, notificationHandler)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		logger.Info("Starting server", "port", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("listen failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
	}
	<-ctx.Done()
	logger.Info("Server exiting")
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
