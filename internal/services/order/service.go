package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/vageeshabr/order-service/internal/filters"
	"github.com/vageeshabr/order-service/internal/models"
	"github.com/vageeshabr/order-service/internal/stores"
)

type service struct {
	store stores.OrderStorer
}

func New(store stores.OrderStorer) *service {
	return &service{store: store}
}

type InvalidParam struct {
	Param string
}

func (i InvalidParam) Error() string {
	return fmt.Sprintf("param '%s' is invalid", i.Param)
}

func (s *service) Find(ctx context.Context, customerId int, date string) ([]*models.Order, error) {
	if customerId <= 0 {
		return nil, InvalidParam{Param: "customerId"}
	}
	if date != "" {
		// validate the date and the format
	}

	res, err := s.store.Find(ctx, filters.Order{Date: date, CustomerId: customerId})
	if err != nil {
		return nil, errors.New("unable to fetch from database")
	}
	return res, nil
}

func (s *service) Create(ctx context.Context, o *stores.OrderCreate) (*models.Order, error) {
	return nil, nil
}
