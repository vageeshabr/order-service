package stores

import (
	"context"
	"github.com/vageeshabr/order-service/internal/filters"
	"github.com/vageeshabr/order-service/internal/models"
)

type OrderStorer interface {
	Find(ctx context.Context, f filters.Order) ([]*models.Order, error)
	Create(ctx context.Context, o *OrderCreate) (*models.Order, error)
}
