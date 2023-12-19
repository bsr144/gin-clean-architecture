package response

type (
	LoginUser struct {
		AccessToken string `json:"access_token"`
	}

	CreateUser struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}
)
