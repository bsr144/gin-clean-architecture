package request

type (
	LoginUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
