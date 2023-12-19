package request

type (
	CreateOrder struct {
		CustomerID  int     `json:"customer_id"`
		OrderDate   string  `json:"order_date"`
		Status      string  `json:"status"`
		TotalAmount float64 `json:"total_amount"`
	}

	UpdateOrder struct {
		ID          int
		CustomerID  int     `json:"customer_id"`
		OrderDate   string  `json:"order_date"`
		Status      string  `json:"status"`
		TotalAmount float64 `json:"total_amount"`
	}

	DeleteOrder struct {
		ID int
	}

	GetOrder struct {
		ID int
	}
)
