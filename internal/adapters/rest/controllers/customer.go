package controllers

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/helpers"
	"dbo-be-task/internal/usecases/customer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	ErrorHelper     *helpers.ErrorHelper
	ParserHelper    *helpers.ParserHelper
	Log             *logrus.Logger
	CustomerUsecase customer.CustomerUseCase
}

func NewCustomerController(errorHelper *helpers.ErrorHelper, parserHelper *helpers.ParserHelper, log *logrus.Logger, customerUsecase customer.CustomerUseCase) *CustomerController {
	return &CustomerController{
		ErrorHelper:     errorHelper,
		ParserHelper:    parserHelper,
		Log:             log,
		CustomerUsecase: customerUsecase,
	}
}

func (c *CustomerController) UpsertCustomer(ctx *gin.Context) {
	var (
		upsertCustomerRequest *request.UpsertCustomer
		dboError              *helpers.Error
	)

	dboError = c.ParserHelper.BindJSON(ctx, &upsertCustomerRequest)

	if dboError != nil {
		return
	}

	upsertCustomerRequest.UserID = c.ParserHelper.GetIntCtx(ctx, "user_id")
	if upsertCustomerRequest.UserID == 0 {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseCreateCustomer, dboError := c.CustomerUsecase.UpsertCustomer(context, upsertCustomerRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusCreated, response.NewHTTPResponseSuccess(http.StatusCreated, "customer successfully registered", responseCreateCustomer))
}

func (c *CustomerController) UpdateCustomerByID(ctx *gin.Context) {
	var (
		updateCustomerRequest *request.UpdateCustomer
		dboError              *helpers.Error
	)

	dboError = c.ParserHelper.BindJSON(ctx, &updateCustomerRequest)

	if dboError != nil {
		return
	}

	updateCustomerRequest.ID = c.ParserHelper.GetIntParam(ctx, "id")

	if updateCustomerRequest.ID == 0 {
		return
	}

	updateCustomerRequest.UserID = c.ParserHelper.GetIntCtx(ctx, "user_id")

	if updateCustomerRequest.UserID == 0 {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseUpdateCustomer, dboError := c.CustomerUsecase.UpdateCustomerByID(context, updateCustomerRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccess(http.StatusOK, "customer successfully updated", responseUpdateCustomer))
}

func (c *CustomerController) DeleteCustomerByID(ctx *gin.Context) {
	customerDeleteRequest := new(request.DeleteCustomer)
	dboError := new(helpers.Error)

	customerDeleteRequest.ID = c.ParserHelper.GetIntParam(ctx, "id")

	if customerDeleteRequest.ID == 0 {
		return
	}

	customerDeleteRequest.UserID = c.ParserHelper.GetIntCtx(ctx, "user_id")

	if customerDeleteRequest.UserID == 0 {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	dboError = c.CustomerUsecase.DeleteCustomerByID(context, customerDeleteRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccess(http.StatusOK, "customer successfully deleted", nil))
}

func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	customerGetByIDRequest := new(request.GetCustomer)
	dboError := new(helpers.Error)

	customerGetByIDRequest.ID = c.ParserHelper.GetIntParam(ctx, "id")

	if customerGetByIDRequest.ID == 0 {
		return
	}

	customerGetByIDRequest.UserID = c.ParserHelper.GetIntCtx(ctx, "user_id")

	if customerGetByIDRequest.UserID == 0 {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseGetCustomer, dboError := c.CustomerUsecase.GetCustomerByID(context, customerGetByIDRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccess(http.StatusOK, "successfully get customer", responseGetCustomer))
}

func (c *CustomerController) GetCustomers(ctx *gin.Context) {
	queryParams := new(request.Query)

	dboError := c.ParserHelper.BindQueryParams(ctx, queryParams)

	if dboError != nil {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseGetCustomers, responsePagination, dboError := c.CustomerUsecase.GetCustomers(context, queryParams)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccessWithPagination(http.StatusOK, "successfully get all customers", responseGetCustomers, responsePagination))
}
