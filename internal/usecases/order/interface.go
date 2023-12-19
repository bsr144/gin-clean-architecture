package order

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/helpers"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, *helpers.Error)
	GetOrders(ctx context.Context, queryParamsRequest *request.Query) ([]*entities.Order, int, *helpers.Error)
	GetOrderByID(ctx context.Context, ID int) (*entities.Order, *helpers.Error)
	UpdateOrder(ctx context.Context, order *entities.Order) (*entities.Order, *helpers.Error)
	DeleteOrderByID(ctx context.Context, ID int) *helpers.Error
}

type OrderPresenter interface {
	PresentCreateOrder(entityOrder *entities.Order) *response.UpsertOrder
	PresentUpdateOrder(entityOrder *entities.Order) *response.UpsertOrder
	PresentGetOrder(entityOrder *entities.Order) *response.GetOrder
	PresentGetOrders(entityOrder []*entities.Order, queryParamsRequest *request.Query, count int) ([]*response.GetOrder, *response.PaginationResponse)
}

type OrderUseCase interface {
	CreateOrder(context context.Context, orderCreateRequest *request.CreateOrder) (*response.UpsertOrder, *helpers.Error)
	GetOrders(context context.Context, queryParamsRequest *request.Query) ([]*response.GetOrder, *response.PaginationResponse, *helpers.Error)
	GetOrderByID(context context.Context, orderUpdateRequest *request.GetOrder) (*response.GetOrder, *helpers.Error)
	UpdateOrderByID(context context.Context, orderUpdateRequest *request.UpdateOrder) (*response.UpsertOrder, *helpers.Error)
	DeleteOrderByID(context context.Context, orderDeleteRequest *request.DeleteOrder) *helpers.Error
}
