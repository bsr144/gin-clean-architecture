package repositories

import (
	"context"
	"database/sql"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/helpers"
	"dbo-be-task/internal/usecases/order"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type orderRepository struct {
	ErrorHelper *helpers.ErrorHelper
	DB          *sql.DB
	Log         *logrus.Logger
}

func NewOrderRepository(errorHelper *helpers.ErrorHelper, db *sql.DB, log *logrus.Logger) order.OrderRepository {
	return &orderRepository{
		ErrorHelper: errorHelper,
		DB:          db,
		Log:         log,
	}
}

func (r *orderRepository) CreateOrder(ctx context.Context, entityOrder *entities.Order) (*entities.Order, *helpers.Error) {
	err := r.DB.
		QueryRowContext(ctx, CREATE_NEW_ORDER_WITH_RETURNING_QUERY, entityOrder.CustomerID, entityOrder.OrderDate, entityOrder.Status, entityOrder.TotalAmount).
		Scan(&entityOrder.ID,
			&entityOrder.CustomerID,
			&entityOrder.OrderDate,
			&entityOrder.Status,
			&entityOrder.TotalAmount,
			&entityOrder.CreatedAt,
			&entityOrder.UpdatedAt,
			&entityOrder.DeletedAt,
		)

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while create order on database", err.Error())
	}

	return entityOrder, nil
}

func (r *orderRepository) GetOrders(ctx context.Context, queryparamRequest *request.Query) ([]*entities.Order, int, *helpers.Error) {
	var (
		orders []*entities.Order
		count  int
	)

	baseQuery := GET_ORDERS_QUERY_PAGINATION
	queryParams := []interface{}{}

	if queryparamRequest.Search != "" {
		baseQuery += WHERE_NAME_ILIKE_PAGINATION
		queryParams = append(queryParams, "%"+queryparamRequest.Search+"%")
	}

	offset := (queryparamRequest.Page - 1) * queryparamRequest.Size
	baseQuery += fmt.Sprintf(PAGINATION_QUERY, offset, queryparamRequest.Size)

	rowStructures, err := r.DB.QueryContext(ctx, baseQuery, queryParams...)
	if err != nil {
		return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while fetching orders", err.Error())
	}
	defer rowStructures.Close()

	for rowStructures.Next() {
		var order entities.Order
		if err := rowStructures.Scan(
			&order.ID,
			&order.CustomerID,
			&order.OrderDate,
			&order.Status,
			&order.TotalAmount,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.DeletedAt,
		); err != nil {
			return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while scanning each rowStructure of orders", err.Error())
		}
		orders = append(orders, &order)
	}

	if err := rowStructures.Err(); err != nil {
		return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "rowStructures contain error", err.Error())
	}

	count, err = r.getCountOfOrders(ctx)
	if err != nil {
		return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while counting orders", err.Error())
	}

	return orders, count, nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, entityOrder *entities.Order) (*entities.Order, *helpers.Error) {
	fmt.Println(entityOrder)
	err := r.DB.
		QueryRowContext(ctx, UPDATE_ORDER_WITH_RETURNING_QUERY, &entityOrder.CustomerID, &entityOrder.OrderDate, &entityOrder.Status, &entityOrder.TotalAmount, &entityOrder.ID).
		Scan(&entityOrder.ID,
			&entityOrder.CustomerID,
			&entityOrder.OrderDate,
			&entityOrder.Status,
			&entityOrder.TotalAmount,
			&entityOrder.CreatedAt,
			&entityOrder.UpdatedAt,
			&entityOrder.DeletedAt,
		)

	if err == sql.ErrNoRows {
		return entityOrder, nil
	}

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while update order data on database", err.Error())
	}

	return entityOrder, nil
}

func (r *orderRepository) getCountOfOrders(ctx context.Context) (int, error) {
	var count int

	err := r.DB.QueryRowContext(ctx, COUNT_ORDERS_QUERY).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *orderRepository) GetOrderByID(ctx context.Context, ID int) (*entities.Order, *helpers.Error) {
	entityOrder := new(entities.Order)

	err := r.DB.
		QueryRowContext(ctx, GET_ORDER_BY_ID_QUERY, ID).
		Scan(&entityOrder.ID,
			&entityOrder.CustomerID,
			&entityOrder.OrderDate,
			&entityOrder.Status,
			&entityOrder.TotalAmount,
			&entityOrder.CreatedAt,
			&entityOrder.UpdatedAt,
			&entityOrder.DeletedAt,
		)

	if err == sql.ErrNoRows {
		return nil, r.ErrorHelper.NewError(http.StatusNotFound, "order not found", sql.ErrNoRows.Error())
	}

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while get order by id from database", err.Error())
	}

	return entityOrder, nil
}

func (r *orderRepository) DeleteOrderByID(ctx context.Context, ID int) *helpers.Error {
	_, err := r.DB.ExecContext(ctx, DELETE_ORDER_QUERY, ID)

	if err != nil {
		return r.ErrorHelper.NewError(http.StatusInternalServerError, "error while deleting order by id from database", err.Error())
	}

	return nil
}
