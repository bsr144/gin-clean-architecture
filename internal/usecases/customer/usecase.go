package customer

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/helpers"
	"net/http"

	"github.com/sirupsen/logrus"
)

type usecase struct {
	CustomerRepository CustomerRepository
	CustomerPresenter  CustomerPresenter
	ErrorHelper        *helpers.ErrorHelper
	Log                *logrus.Logger
}

func NewCustomerUsecase(repository CustomerRepository, presenter CustomerPresenter, errorHelper *helpers.ErrorHelper, log *logrus.Logger) CustomerUseCase {
	return &usecase{
		CustomerRepository: repository,
		CustomerPresenter:  presenter,
		ErrorHelper:        errorHelper,
		Log:                log,
	}
}

func (u *usecase) UpsertCustomer(context context.Context, customerUpsertRequest *request.UpsertCustomer) (*response.UpsertCustomer, *helpers.Error) {
	customer, dboError := u.CustomerRepository.GetCustomerByUserID(context, customerUpsertRequest.UserID)

	if dboError != nil {
		return nil, dboError
	}

	// Check if exist, if already exist, we update the customer data
	if customer.IsExist() && customer.IsDeleted() {
		return u.upsertDeletedCustomer(context, customer, customerUpsertRequest)
	} else if customer.IsExist() {
		return u.upsertUpdateCustomer(context, customer, customerUpsertRequest)
	}

	// Else, we create new customer
	return u.upsertCreateCustomer(context, customer, customerUpsertRequest)
}

func (u *usecase) UpdateCustomerByID(context context.Context, customerUpdateRequest *request.UpdateCustomer) (*response.UpsertCustomer, *helpers.Error) {
	customer, dboError := u.CustomerRepository.GetCustomerByID(context, customerUpdateRequest.ID)

	if dboError != nil {
		return nil, dboError
	}

	if !customer.IsEqualRequestedUserID(customerUpdateRequest.UserID) {
		return nil, u.ErrorHelper.NewError(http.StatusUnauthorized, "you are unauthorized to do this request", "you are unauthorized to do this request")
	}

	return u.updateCustomer(context, customer, customerUpdateRequest)
}

func (u *usecase) DeleteCustomerByID(context context.Context, customerDeleteRequest *request.DeleteCustomer) *helpers.Error {
	customer, dboError := u.CustomerRepository.GetCustomerByID(context, customerDeleteRequest.ID)

	if dboError != nil {
		return dboError
	}

	if !customer.IsEqualRequestedUserID(customerDeleteRequest.UserID) {
		return u.ErrorHelper.NewError(http.StatusUnauthorized, "you are unauthorized to do this request", "you are unauthorized to do this request")
	}

	return u.CustomerRepository.DeleteCustomerByID(context, customerDeleteRequest.ID)
}

func (u *usecase) GetCustomerByID(context context.Context, customerGetByIDRequest *request.GetCustomer) (*response.GetCustomer, *helpers.Error) {
	customer, dboError := u.CustomerRepository.GetCustomerByID(context, customerGetByIDRequest.ID)

	if dboError != nil {
		return nil, dboError
	}

	if !customer.IsEqualRequestedUserID(customerGetByIDRequest.UserID) {
		return nil, u.ErrorHelper.NewError(http.StatusUnauthorized, "you are unauthorized to do this request", "you are unauthorized to do this request")
	}

	return u.CustomerPresenter.PresentGetCustomer(customer), nil
}

func (u *usecase) GetCustomers(context context.Context, queryParamsRequest *request.Query) ([]*response.GetCustomer, *response.PaginationResponse, *helpers.Error) {
	customers, count, dboError := u.CustomerRepository.GetCustomers(context, queryParamsRequest)

	if dboError != nil {
		return nil, nil, dboError
	}

	responseGetCustomers, responsePagination := u.CustomerPresenter.PresentGetCustomers(customers, queryParamsRequest, count)

	return responseGetCustomers, responsePagination, nil
}

func (u *usecase) upsertCreateCustomer(context context.Context, customer *entities.Customer, customerUpsertRequest *request.UpsertCustomer) (*response.UpsertCustomer, *helpers.Error) {
	customer.SetCreateDataUpsert(customerUpsertRequest)

	customer, dboError := u.CustomerRepository.CreateCustomer(context, customer)

	if dboError != nil {
		return nil, dboError
	}

	return u.CustomerPresenter.PresentUpsertCustomer(customer), nil
}

func (u *usecase) upsertUpdateCustomer(context context.Context, customer *entities.Customer, customerUpsertRequest *request.UpsertCustomer) (*response.UpsertCustomer, *helpers.Error) {
	customer.SetUpdateDataUpsert(customerUpsertRequest)

	customer, dboError := u.CustomerRepository.UpdateCustomer(context, customer)

	if dboError != nil {
		return nil, dboError
	}

	return u.CustomerPresenter.PresentUpsertCustomer(customer), nil
}

func (u *usecase) upsertDeletedCustomer(context context.Context, customer *entities.Customer, customerUpsertRequest *request.UpsertCustomer) (*response.UpsertCustomer, *helpers.Error) {
	customer.SetUpdateDataUpsert(customerUpsertRequest)

	customer, dboError := u.CustomerRepository.UpdateDeletedCustomer(context, customer)

	if dboError != nil {
		return nil, dboError
	}

	return u.CustomerPresenter.PresentUpsertCustomer(customer), nil
}

func (u *usecase) updateCustomer(context context.Context, customer *entities.Customer, customerUpdateRequest *request.UpdateCustomer) (*response.UpsertCustomer, *helpers.Error) {
	customer.SetUpdateData(customerUpdateRequest)

	customer, dboError := u.CustomerRepository.UpdateCustomer(context, customer)

	if dboError != nil {
		return nil, dboError
	}

	return u.CustomerPresenter.PresentUpsertCustomer(customer), nil
}
