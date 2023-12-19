package repositories

const (
	WHERE_NAME_ILIKE_PAGINATION = " AND name ILIKE $1"
	PAGINATION_QUERY            = " OFFSET %d LIMIT %d;"
)
