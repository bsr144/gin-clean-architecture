package customer

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/helpers"
)

type CustomerRepository interface {
	GetCustomers(ctx context.Context, queryParamsRequest *request.Query) ([]*entities.Customer, int, *helpers.Error)
	CreateCustomer(ctx context.Context, customer *entities.Customer) (*entities.Customer, *helpers.Error)
	UpdateCustomer(ctx context.Context, customer *entities.Customer) (*entities.Customer, *helpers.Error)
	GetCustomerByID(ctx context.Context, ID int) (*entities.Customer, *helpers.Error)
	DeleteCustomerByID(ctx context.Context, ID int) *helpers.Error
	GetCustomerByUserID(ctx context.Context, userID int) (*entities.Customer, *helpers.Error)
	UpdateDeletedCustomer(ctx context.Context, customer *entities.Customer) (*entities.Customer, *helpers.Error)
}

type CustomerPresenter interface {
	PresentUpsertCustomer(entityCustomer *entities.Customer) *response.UpsertCustomer
	PresentGetCustomer(entityCustomer *entities.Customer) *response.GetCustomer
	PresentGetCustomers(entityCustomer []*entities.Customer, queryParamsRequest *request.Query, count int) ([]*response.GetCustomer, *response.PaginationResponse)
}

type CustomerUseCase interface {
	UpsertCustomer(context context.Context, customerUpsertRequest *request.UpsertCustomer) (*response.UpsertCustomer, *helpers.Error)
	UpdateCustomerByID(context context.Context, customerUpdateRequest *request.UpdateCustomer) (*response.UpsertCustomer, *helpers.Error)
	DeleteCustomerByID(context context.Context, customerDeleteRequest *request.DeleteCustomer) *helpers.Error
	GetCustomerByID(context context.Context, customerUpdateRequest *request.GetCustomer) (*response.GetCustomer, *helpers.Error)
	GetCustomers(context context.Context, queryParamsRequest *request.Query) ([]*response.GetCustomer, *response.PaginationResponse, *helpers.Error)
}
