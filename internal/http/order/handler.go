package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vageeshabr/order-service/internal/models"
	"github.com/vageeshabr/order-service/internal/services"
	"net/http"
	"strconv"
	"time"
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

func (h *handler) GET(w http.ResponseWriter, r *http.Request) {
	customerId := 0
	if v := r.URL.Query().Get("customerId"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"param customerId is not int"}`))
			return
		}
		customerId = id
	}

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	orders, err := h.svc.Find(ctx, customerId, r.URL.Query().Get("date"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

	res := listResponse{
		Data: struct {
			Orders []*models.Order `json:"orders"`
		}{
			Orders: orders,
		},
	}

	json.NewEncoder(w).Encode(res)
}

func (h *handler) POST(w http.ResponseWriter, r *http.Request) {

}
