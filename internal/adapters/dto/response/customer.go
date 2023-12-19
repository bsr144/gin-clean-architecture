package response

type (
	GetCustomer struct {
		ID      int    `json:"id"`
		UserID  int    `json:"user_id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}
	UpsertCustomer struct {
		ID      int    `json:"id"`
		UserID  int    `json:"user_id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}
)
