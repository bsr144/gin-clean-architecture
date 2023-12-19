package repositories

import (
	"context"
	"database/sql"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/helpers"
	"dbo-be-task/internal/usecases/customer"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type customerRepository struct {
	ErrorHelper *helpers.ErrorHelper
	DB          *sql.DB
	Log         *logrus.Logger
}

func NewCustomerRepository(errorHelper *helpers.ErrorHelper, db *sql.DB, log *logrus.Logger) customer.CustomerRepository {
	return &customerRepository{
		ErrorHelper: errorHelper,
		DB:          db,
		Log:         log,
	}
}

func (r *customerRepository) CreateCustomer(ctx context.Context, entityCustomer *entities.Customer) (*entities.Customer, *helpers.Error) {
	err := r.DB.
		QueryRowContext(ctx, CREATE_NEW_CUSTOMER_WITH_RETURNING_QUERY, entityCustomer.UserID, entityCustomer.Name, entityCustomer.Phone, entityCustomer.Address).
		Scan(&entityCustomer.ID,
			&entityCustomer.UserID,
			&entityCustomer.Name,
			&entityCustomer.Phone,
			&entityCustomer.Address,
			&entityCustomer.CreatedAt,
			&entityCustomer.UpdatedAt,
			&entityCustomer.DeletedAt,
		)

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while create customer on database", err.Error())
	}

	return entityCustomer, nil
}

func (r *customerRepository) UpdateCustomer(ctx context.Context, entityCustomer *entities.Customer) (*entities.Customer, *helpers.Error) {
	err := r.DB.
		QueryRowContext(ctx, UPDATE_CUSTOMER_WITH_RETURNING_QUERY, &entityCustomer.Name, &entityCustomer.Phone, &entityCustomer.Address, &entityCustomer.ID).
		Scan(&entityCustomer.ID,
			&entityCustomer.UserID,
			&entityCustomer.Name,
			&entityCustomer.Phone,
			&entityCustomer.Address,
			&entityCustomer.CreatedAt,
			&entityCustomer.UpdatedAt,
			&entityCustomer.DeletedAt,
		)

	if err == sql.ErrNoRows {
		return entityCustomer, nil
	}

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while update customer data on database", err.Error())
	}

	return entityCustomer, nil
}

func (r *customerRepository) UpdateDeletedCustomer(ctx context.Context, entityCustomer *entities.Customer) (*entities.Customer, *helpers.Error) {
	err := r.DB.
		QueryRowContext(ctx, UPDATE_DELETED_CUSTOMER_WITH_RETURNING_QUERY, &entityCustomer.Name, &entityCustomer.Phone, &entityCustomer.Address, &entityCustomer.ID).
		Scan(&entityCustomer.ID,
			&entityCustomer.UserID,
			&entityCustomer.Name,
			&entityCustomer.Phone,
			&entityCustomer.Address,
			&entityCustomer.CreatedAt,
			&entityCustomer.UpdatedAt,
			&entityCustomer.DeletedAt,
		)

	if err == sql.ErrNoRows {
		return entityCustomer, nil
	}

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while update customer data on database", err.Error())
	}

	return entityCustomer, nil
}

func (r *customerRepository) GetCustomers(ctx context.Context, queryparamRequest *request.Query) ([]*entities.Customer, int, *helpers.Error) {
	var (
		customers []*entities.Customer
		count     int
	)

	baseQuery := GET_CUSTOMERS_QUERY_PAGINATION
	queryParams := []interface{}{}

	if queryparamRequest.Search != "" {
		baseQuery += WHERE_NAME_ILIKE_PAGINATION
		queryParams = append(queryParams, "%"+queryparamRequest.Search+"%")
	}

	offset := (queryparamRequest.Page - 1) * queryparamRequest.Size
	baseQuery += fmt.Sprintf(PAGINATION_QUERY, offset, queryparamRequest.Size)

	fmt.Println(baseQuery, queryParams)

	rowStructures, err := r.DB.QueryContext(ctx, baseQuery, queryParams...)
	if err != nil {
		return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while fetching customers", err.Error())
	}
	defer rowStructures.Close()

	for rowStructures.Next() {
		var customer entities.Customer
		if err := rowStructures.Scan(
			&customer.ID,
			&customer.UserID,
			&customer.Name,
			&customer.Phone,
			&customer.Address,
			&customer.CreatedAt,
			&customer.UpdatedAt,
			&customer.DeletedAt,
		); err != nil {
			return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while scanning each rowStructure of customers", err.Error())
		}
		customers = append(customers, &customer)
	}

	if err := rowStructures.Err(); err != nil {
		return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "rowStructures contain error", err.Error())
	}

	count, err = r.getCountOfCustomers(ctx)
	if err != nil {
		return nil, count, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while counting customers", err.Error())
	}

	return customers, count, nil
}

func (r *customerRepository) getCountOfCustomers(ctx context.Context) (int, error) {
	var count int

	err := r.DB.QueryRowContext(ctx, COUNT_CUSTOMERS_QUERY).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *customerRepository) GetCustomerByUserID(ctx context.Context, userID int) (*entities.Customer, *helpers.Error) {
	entityCustomer := new(entities.Customer)

	err := r.DB.
		QueryRowContext(ctx, GET_CUSTOMER_BY_USER_ID_QUERY, userID).
		Scan(&entityCustomer.ID,
			&entityCustomer.UserID,
			&entityCustomer.Name,
			&entityCustomer.Phone,
			&entityCustomer.Address,
			&entityCustomer.CreatedAt,
			&entityCustomer.UpdatedAt,
			&entityCustomer.DeletedAt,
		)

	if err == sql.ErrNoRows {
		return entityCustomer, nil
	}

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while get customer by user_id from database", err.Error())
	}

	return entityCustomer, nil
}

func (r *customerRepository) GetCustomerByID(ctx context.Context, ID int) (*entities.Customer, *helpers.Error) {
	entityCustomer := new(entities.Customer)

	err := r.DB.
		QueryRowContext(ctx, GET_CUSTOMER_BY_ID_QUERY, ID).
		Scan(&entityCustomer.ID,
			&entityCustomer.UserID,
			&entityCustomer.Name,
			&entityCustomer.Phone,
			&entityCustomer.Address,
			&entityCustomer.CreatedAt,
			&entityCustomer.UpdatedAt,
			&entityCustomer.DeletedAt,
		)

	if err == sql.ErrNoRows {
		return nil, r.ErrorHelper.NewError(http.StatusNotFound, "customer not found", sql.ErrNoRows.Error())
	}

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while get customer by id from database", err.Error())
	}

	return entityCustomer, nil
}

func (r *customerRepository) DeleteCustomerByID(ctx context.Context, ID int) *helpers.Error {
	_, err := r.DB.ExecContext(ctx, DELETE_CUSTOMER_QUERY, ID)

	if err != nil {
		return r.ErrorHelper.NewError(http.StatusInternalServerError, "error while deleting customer by id from database", err.Error())
	}

	return nil
}
