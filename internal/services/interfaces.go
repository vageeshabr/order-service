package services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/vageeshabr/order-service/internal/models"
	"github.com/vageeshabr/order-service/internal/stores"
)

type Order interface {
	Find(ctx gofr.Context, customerId int, date string) ([]*models.Order, error)
	Create(ctx gofr.Context, o *stores.OrderCreate) (*models.Order, error)
}
