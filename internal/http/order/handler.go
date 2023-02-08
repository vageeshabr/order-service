package order

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/types"
	"github.com/vageeshabr/order-service/internal/models"
	"github.com/vageeshabr/order-service/internal/services"
	"strconv"
)

type handler struct {
	svc services.Order
}

func New(svc services.Order) *handler {
	return &handler{svc: svc}
}

type listResponse struct {
	Data struct {
		Orders []*models.Order `json:"orders"`
	} `json:"data"`
}

func (h *handler) Index(c *gofr.Context) (interface{}, error) {
	customerId := 0

	if v := c.Param("customerId"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.InvalidParam{Param: []string{"customerId"}}
		}
		customerId = id
	}
	return types.Response{Data: customerId}, nil
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {

	return nil, nil

}

func (h *handler) Delete(ctx *gofr.Context) (interface{}, error) {

	return nil, nil

}
