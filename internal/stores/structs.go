package stores

import (
	"encoding/json"
	"time"
)

type OrderCreate struct {
	CustomerID int             `json:"customerId"`
	Address    string          `json:"address"`
	Cart       json.RawMessage `json:"cart"`
	Date       time.Time       `json:"date"`
}
