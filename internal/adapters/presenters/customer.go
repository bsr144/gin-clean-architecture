package presenters

import (
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/usecases/customer"

	"github.com/sirupsen/logrus"
)

type customerPresenter struct {
	Log *logrus.Logger
}

func NewCustomerPresenter(log *logrus.Logger) customer.CustomerPresenter {
	return &customerPresenter{
		Log: log,
	}
}

func (p *customerPresenter) PresentUpsertCustomer(entityCustomer *entities.Customer) *response.UpsertCustomer {
	return &response.UpsertCustomer{
		ID:      entityCustomer.ID,
		UserID:  entityCustomer.UserID,
		Name:    entityCustomer.Name,
		Phone:   entityCustomer.Phone,
		Address: entityCustomer.Address,
	}
}

func (p *customerPresenter) PresentGetCustomer(entityCustomer *entities.Customer) *response.GetCustomer {
	return &response.GetCustomer{
		ID:      entityCustomer.ID,
		UserID:  entityCustomer.UserID,
		Name:    entityCustomer.Name,
		Phone:   entityCustomer.Phone,
		Address: entityCustomer.Address,
	}
}

func (p *customerPresenter) PresentGetCustomers(entityCustomers []*entities.Customer, queryParamsRequest *request.Query, count int) ([]*response.GetCustomer, *response.PaginationResponse) {
	presentedCustomers := make([]*response.GetCustomer, 0, len(entityCustomers))

	for _, entityCustomer := range entityCustomers {
		presentedCustomer := &response.GetCustomer{
			ID:      entityCustomer.ID,
			UserID:  entityCustomer.UserID,
			Name:    entityCustomer.Name,
			Phone:   entityCustomer.Phone,
			Address: entityCustomer.Address,
		}

		presentedCustomers = append(presentedCustomers, presentedCustomer)
	}

	pagination := calculatePagination(count, queryParamsRequest.Page, queryParamsRequest.Size, len(presentedCustomers))

	return presentedCustomers, pagination
}
