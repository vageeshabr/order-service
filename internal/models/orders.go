package models

import (
	"encoding/json"
	"time"
)

type Order struct {
	Id         int             `json:"id"`
	CustomerID int             `json:"customerId"`
	Address    string          `json:"address"`
	Cart       json.RawMessage `json:"cart"`
	Date       time.Time       `json:"date"`
	Amount     float64         `json:"amount"`
}
