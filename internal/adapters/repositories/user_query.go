package repositories

const (
	// Write queries
	CREATE_NEW_USER_WITH_RETURNING_QUERY = "INSERT INTO users (email, password, salt) VALUES ($1, $2, $3) RETURNING id, email;"

	// Read queries
	GET_USER_BY_EMAIL_QUERY = "SELECT * FROM users WHERE email = ($1) AND deleted_at IS NULL;"
)
