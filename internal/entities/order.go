package entities

import (
	"dbo-be-task/internal/adapters/dto/request"
	"time"
)

type Order struct {
	ID          int
	CustomerID  int
	OrderDate   time.Time
	Status      string
	TotalAmount float64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

func (o *Order) SetCreateData(orderCreateRequest *request.CreateOrder) error {
	orderDateTime, err := time.Parse("2006-01-02", orderCreateRequest.OrderDate)

	o.CustomerID = orderCreateRequest.CustomerID
	o.OrderDate = orderDateTime
	o.Status = orderCreateRequest.Status
	o.TotalAmount = orderCreateRequest.TotalAmount

	if err != nil {
		return err
	}

	return nil
}

func (o *Order) SetUpdateData(orderUpdateRequest *request.UpdateOrder) error {
	orderDateTime, err := time.Parse("2006-01-02", orderUpdateRequest.OrderDate)

	o.ID = orderUpdateRequest.ID
	o.CustomerID = orderUpdateRequest.CustomerID
	o.OrderDate = orderDateTime
	o.Status = orderUpdateRequest.Status
	o.TotalAmount = orderUpdateRequest.TotalAmount

	if err != nil {
		return err
	}

	return nil
}
