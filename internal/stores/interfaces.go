package stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/vageeshabr/order-service/internal/filters"
	"github.com/vageeshabr/order-service/internal/models"
)

type OrderStorer interface {
	Find(ctx gofr.Context, f filters.Order) ([]*models.Order, error)
	Create(ctx gofr.Context, o *OrderCreate) (*models.Order, error)
}
