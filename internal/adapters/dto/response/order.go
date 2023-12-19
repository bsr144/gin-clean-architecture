package response

import "time"

type (
	GetOrder struct {
		ID          int       `json:"id"`
		CustomerID  int       `json:"customer_id"`
		OrderDate   time.Time `json:"order_date"`
		Status      string    `json:"status"`
		TotalAmount float64   `json:"total_amount"`
		CreatedAt   string    `json:"created_at"`
	}
	UpsertOrder struct {
		ID          int       `json:"id"`
		CustomerID  int       `json:"customer_id"`
		OrderDate   time.Time `json:"order_date"`
		Status      string    `json:"status"`
		TotalAmount float64   `json:"total_amount"`
	}
)
