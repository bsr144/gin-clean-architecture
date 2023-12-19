package entities

import (
	"dbo-be-task/internal/adapters/dto/request"
	"time"
)

type Customer struct {
	ID        int
	UserID    int
	Name      string
	Phone     string
	Address   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (c *Customer) IsExist() bool {
	return c.ID != 0
}

func (c *Customer) IsDeleted() bool {
	return c.DeletedAt != nil
}

func (c *Customer) IsEqualRequestedUserID(requestUserID int) bool {
	return c.UserID == requestUserID
}

func (c *Customer) SetUpdateDataUpsert(customerUpsertRequest *request.UpsertCustomer) {
	c.Name = customerUpsertRequest.Name
	c.Phone = customerUpsertRequest.Phone
	c.Address = customerUpsertRequest.Address
}

func (c *Customer) SetCreateDataUpsert(customerUpsertRequest *request.UpsertCustomer) {
	c.UserID = customerUpsertRequest.UserID
	c.Name = customerUpsertRequest.Name
	c.Phone = customerUpsertRequest.Phone
	c.Address = customerUpsertRequest.Address
}

func (c *Customer) SetUpdateData(customerUpdateRequest *request.UpdateCustomer) {
	c.UserID = customerUpdateRequest.UserID
	c.Name = customerUpdateRequest.Name
	c.Phone = customerUpdateRequest.Phone
	c.Address = customerUpdateRequest.Address
}
