package handler

import (
	"context"
	"net/http"
	"service-customer/internal/dto"
	"service-customer/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type NotificationService interface {
	GetNotifications(ctx context.Context, customerID int64, from, to time.Time, page, size int) ([]model.Notification, int, error)
}

type NotificationHandler struct {
	notificationService NotificationService
}

func NewNotificationHandler(notificationService NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notificationService}
}

func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	customerID, _ := strconv.ParseInt(c.Query("customerId"), 10, 64)
	if customerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customerId is required"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	fromStr := c.Query("from")
	toStr := c.Query("to")

	now := time.Now()
	from := now.AddDate(0, -1, 0)
	to := now

	if fromStr != "" {
		if t, err := time.Parse(time.RFC3339, fromStr); err == nil {
			from = t
		}
	}
	if toStr != "" {
		if t, err := time.Parse(time.RFC3339, toStr); err == nil {
			to = t
		}
	}

	notifications, total, err := h.notificationService.GetNotifications(c.Request.Context(), customerID, from, to, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dtos := make([]*dto.NotificationDTO, len(notifications))
	for i, n := range notifications {
		d := dto.NotificationDTO{
			ID:       n.ID,
			Text:     n.Text,
			EntityID: n.EntityID,
			CreateAt: n.CreateAt,
		}
		dtos[i] = &d
	}

	c.JSON(http.StatusOK, dto.PagedNotifications{
		Paging: dto.Paging{Page: page, Size: size, Total: total},
		Items:  dtos,
	})
}
