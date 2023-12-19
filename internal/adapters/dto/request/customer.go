package request

type (
	UpsertCustomer struct {
		UserID  int
		Name    string
		Phone   string
		Address string
	}

	UpdateCustomer struct {
		ID      int
		UserID  int
		Name    string
		Phone   string
		Address string
	}

	DeleteCustomer struct {
		ID     int
		UserID int
	}

	GetCustomer struct {
		ID     int
		UserID int
	}
)
