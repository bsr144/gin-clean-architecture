package order

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
	OrderRepository OrderRepository
	OrderPresenter  OrderPresenter
	ErrorHelper     *helpers.ErrorHelper
	Log             *logrus.Logger
}

func NewOrderUsecase(repository OrderRepository, presenter OrderPresenter, errorHelper *helpers.ErrorHelper, log *logrus.Logger) OrderUseCase {
	return &usecase{
		OrderRepository: repository,
		OrderPresenter:  presenter,
		ErrorHelper:     errorHelper,
		Log:             log,
	}
}

func (u *usecase) CreateOrder(context context.Context, orderCreateRequest *request.CreateOrder) (*response.UpsertOrder, *helpers.Error) {
	order := new(entities.Order)

	err := order.SetCreateData(orderCreateRequest)

	if err != nil {
		return nil, u.ErrorHelper.NewError(http.StatusInternalServerError, "error while parsing order_date to time", err.Error())
	}

	order, dboError := u.OrderRepository.CreateOrder(context, order)

	if dboError != nil {
		return nil, dboError
	}

	return u.OrderPresenter.PresentCreateOrder(order), nil
}

func (u *usecase) GetOrders(context context.Context, queryParamsRequest *request.Query) ([]*response.GetOrder, *response.PaginationResponse, *helpers.Error) {
	orders, count, dboError := u.OrderRepository.GetOrders(context, queryParamsRequest)

	if dboError != nil {
		return nil, nil, dboError
	}

	responseGetOrders, responsePagination := u.OrderPresenter.PresentGetOrders(orders, queryParamsRequest, count)

	return responseGetOrders, responsePagination, nil
}

func (u *usecase) GetOrderByID(context context.Context, orderGetByIDRequest *request.GetOrder) (*response.GetOrder, *helpers.Error) {
	order, dboError := u.OrderRepository.GetOrderByID(context, orderGetByIDRequest.ID)

	if dboError != nil {
		return nil, dboError
	}

	return u.OrderPresenter.PresentGetOrder(order), nil
}

func (u *usecase) UpdateOrderByID(context context.Context, orderUpdateRequest *request.UpdateOrder) (*response.UpsertOrder, *helpers.Error) {
	order := new(entities.Order)

	order.SetUpdateData(orderUpdateRequest)

	order, dboError := u.OrderRepository.UpdateOrder(context, order)

	if dboError != nil {
		return nil, dboError
	}

	return u.OrderPresenter.PresentUpdateOrder(order), nil

}

func (u *usecase) DeleteOrderByID(context context.Context, orderDeleteRequest *request.DeleteOrder) *helpers.Error {
	_, dboError := u.OrderRepository.GetOrderByID(context, orderDeleteRequest.ID)

	if dboError != nil {
		return dboError
	}

	return u.OrderRepository.DeleteOrderByID(context, orderDeleteRequest.ID)
}
