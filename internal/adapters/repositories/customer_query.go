package repositories

const (
	// // Write queries
	CREATE_NEW_CUSTOMER_WITH_RETURNING_QUERY = `
		INSERT INTO customers (user_id, name, phone, address) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, user_id, name, phone, address, created_at, updated_at, deleted_at;
	`

	UPDATE_CUSTOMER_WITH_RETURNING_QUERY = `
		UPDATE customers
		SET name = $1, phone = $2, address = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING id, user_id, name, phone, address, created_at, updated_at, deleted_at;
	`

	UPDATE_DELETED_CUSTOMER_WITH_RETURNING_QUERY = `
		UPDATE customers
		SET name = $1, phone = $2, address = $3, created_at = NOW(), updated_at = NOW(), deleted_at = NULL
		WHERE id = $4
		RETURNING id, user_id, name, phone, address, created_at, updated_at, deleted_at;
	`

	DELETE_CUSTOMER_QUERY = `UPDATE customers SET deleted_at = NOW() WHERE id = ($1);`

	// Read queries
	GET_CUSTOMER_BY_USER_ID_QUERY  = "SELECT * FROM customers WHERE user_id = ($1);"
	GET_CUSTOMER_BY_ID_QUERY       = "SELECT * FROM customers WHERE id = ($1) AND deleted_at IS NULL;"
	GET_CUSTOMERS_QUERY            = "SELECT * FROM customers WHERE deleted_at IS NULL;"
	GET_CUSTOMERS_QUERY_PAGINATION = "SELECT * FROM customers WHERE deleted_at IS NULL"

	COUNT_CUSTOMERS_QUERY = "SELECT COUNT(*) FROM customers WHERE deleted_at IS NULL;"
)
