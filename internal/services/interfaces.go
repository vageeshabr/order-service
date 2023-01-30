package services

import (
	"context"
	"github.com/vageeshabr/order-service/internal/models"
	"github.com/vageeshabr/order-service/internal/stores"
)

type Order interface {
	Find(ctx context.Context, customerId int, date string) ([]*models.Order, error)
	Create(ctx context.Context, o *stores.OrderCreate) (*models.Order, error)
}
