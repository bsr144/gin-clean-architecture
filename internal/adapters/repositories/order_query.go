package repositories

const (
	// // Write queries
	CREATE_NEW_ORDER_WITH_RETURNING_QUERY = `
		INSERT INTO orders (customer_id, order_date, status, total_amount)
		VALUES ($1, $2, $3, $4)
		RETURNING id, customer_id, order_date, status, total_amount, created_at, updated_at, deleted_at;
	`

	UPDATE_ORDER_WITH_RETURNING_QUERY = `
		UPDATE orders
		SET customer_id = $1, order_date = $2, status = $3, total_amount = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING id, customer_id, order_date, status, total_amount, created_at, updated_at, deleted_at;
	`

	DELETE_ORDER_QUERY = `UPDATE orders SET deleted_at = NOW() WHERE id = ($1);`

	// Read queries
	GET_ORDER_BY_ID_QUERY       = "SELECT * FROM orders WHERE id = ($1) AND deleted_at IS NULL;"
	GET_ORDERS_QUERY            = "SELECT * FROM orders WHERE deleted_at IS NULL;"
	GET_ORDERS_QUERY_PAGINATION = "SELECT * FROM orders WHERE deleted_at IS NULL"

	COUNT_ORDERS_QUERY = "SELECT COUNT(*) FROM orders WHERE deleted_at IS NULL;"
)
