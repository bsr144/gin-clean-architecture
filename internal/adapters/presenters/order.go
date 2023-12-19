package presenters

import (
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/usecases/order"
	"time"

	"github.com/sirupsen/logrus"
)

type orderPresenter struct {
	Log *logrus.Logger
}

func NewOrderPresenter(log *logrus.Logger) order.OrderPresenter {
	return &orderPresenter{
		Log: log,
	}
}

// Both Create and Update looks redundant at first time, but if there is any different fields that want to be added
// You could add or remove it as you like
// Depending on what you need on both create and update
func (p *orderPresenter) PresentCreateOrder(entityOrder *entities.Order) *response.UpsertOrder {
	return &response.UpsertOrder{
		ID:          entityOrder.ID,
		CustomerID:  entityOrder.CustomerID,
		OrderDate:   entityOrder.OrderDate,
		Status:      entityOrder.Status,
		TotalAmount: entityOrder.TotalAmount,
	}
}

func (p *orderPresenter) PresentUpdateOrder(entityOrder *entities.Order) *response.UpsertOrder {
	return &response.UpsertOrder{
		ID:          entityOrder.ID,
		CustomerID:  entityOrder.CustomerID,
		OrderDate:   entityOrder.OrderDate,
		Status:      entityOrder.Status,
		TotalAmount: entityOrder.TotalAmount,
	}
}

func (p *orderPresenter) PresentGetOrder(entityOrder *entities.Order) *response.GetOrder {
	return &response.GetOrder{
		ID:          entityOrder.ID,
		CustomerID:  entityOrder.CustomerID,
		OrderDate:   entityOrder.OrderDate,
		Status:      entityOrder.Status,
		TotalAmount: entityOrder.TotalAmount,
		CreatedAt:   entityOrder.CreatedAt.Format(time.RFC850),
	}
}

func (p *orderPresenter) PresentGetOrders(entityOrders []*entities.Order, queryParamsRequest *request.Query, count int) ([]*response.GetOrder, *response.PaginationResponse) {
	presentedOrders := make([]*response.GetOrder, 0, len(entityOrders))

	for _, entityOrder := range entityOrders {
		presentedOrder := &response.GetOrder{
			ID:          entityOrder.ID,
			CustomerID:  entityOrder.CustomerID,
			OrderDate:   entityOrder.OrderDate,
			Status:      entityOrder.Status,
			TotalAmount: entityOrder.TotalAmount,
			CreatedAt:   entityOrder.CreatedAt.Format(time.RFC850),
		}

		presentedOrders = append(presentedOrders, presentedOrder)
	}

	pagination := calculatePagination(count, queryParamsRequest.Page, queryParamsRequest.Size, len(presentedOrders))

	return presentedOrders, pagination
}
