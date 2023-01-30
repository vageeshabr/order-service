package order

import (
	"context"
	"database/sql"
	"github.com/vageeshabr/order-service/internal/filters"
	"github.com/vageeshabr/order-service/internal/models"
	"github.com/vageeshabr/order-service/internal/stores"
	"strings"
)

const tableName = "orders"

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) Find(ctx context.Context, f filters.Order) (orders []*models.Order, err error) {

	var where []string
	var args []interface{}

	if f.CustomerId > 0 {
		where = append(where, "customer_id = ?")
		args = append(args, f.CustomerId)
	}
	if f.Date != "" {
		where = append(where, "date = ?")
		args = append(args, f.Date)
	}

	query := "select id, customer_id, address, cart, preferred_date, amount from " + tableName

	if len(where) > 0 {
		query += " where " + strings.Join(where, " AND ")
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var o *models.Order

		err := rows.Scan(&o.Id, &o.CustomerID, &o.Address, &o.Cart, &o.Date, &o.Amount)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return
}

func (s *store) Create(ctx context.Context, o *stores.OrderCreate) (*models.Order, error) {

	return nil, nil
}
