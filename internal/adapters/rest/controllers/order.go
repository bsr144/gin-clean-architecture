package controllers

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/helpers"
	"dbo-be-task/internal/usecases/order"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type OrderController struct {
	ErrorHelper  *helpers.ErrorHelper
	ParserHelper *helpers.ParserHelper
	Log          *logrus.Logger
	OrderUsecase order.OrderUseCase
}

func NewOrderController(errorHelper *helpers.ErrorHelper, parserHelper *helpers.ParserHelper, log *logrus.Logger, orderUsecase order.OrderUseCase) *OrderController {
	return &OrderController{
		ErrorHelper:  errorHelper,
		ParserHelper: parserHelper,
		Log:          log,
		OrderUsecase: orderUsecase,
	}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var (
		createOrderRequest *request.CreateOrder
		dboError           *helpers.Error
	)

	dboError = c.ParserHelper.BindJSON(ctx, &createOrderRequest)

	if dboError != nil {
		return
	}

	fmt.Println(createOrderRequest)

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseCreateOrder, dboError := c.OrderUsecase.CreateOrder(context, createOrderRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusCreated, response.NewHTTPResponseSuccess(http.StatusCreated, "order successfully registered", responseCreateOrder))
}

func (c *OrderController) GetOrders(ctx *gin.Context) {
	queryParams := new(request.Query)

	dboError := c.ParserHelper.BindQueryParams(ctx, queryParams)

	if dboError != nil {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseGetOrders, responsePagination, dboError := c.OrderUsecase.GetOrders(context, queryParams)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccessWithPagination(http.StatusOK, "successfully get all orders", responseGetOrders, responsePagination))
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	orderGetByIDRequest := new(request.GetOrder)
	dboError := new(helpers.Error)

	orderGetByIDRequest.ID = c.ParserHelper.GetIntParam(ctx, "id")

	if orderGetByIDRequest.ID == 0 {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseGetOrder, dboError := c.OrderUsecase.GetOrderByID(context, orderGetByIDRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccess(http.StatusOK, "successfully get order", responseGetOrder))
}

func (c *OrderController) UpdateOrderByID(ctx *gin.Context) {
	var (
		updateOrderRequest *request.UpdateOrder
		dboError           *helpers.Error
	)

	dboError = c.ParserHelper.BindJSON(ctx, &updateOrderRequest)

	if dboError != nil {
		return
	}

	updateOrderRequest.ID = c.ParserHelper.GetIntParam(ctx, "id")

	if updateOrderRequest.ID == 0 {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseUpdateOrder, dboError := c.OrderUsecase.UpdateOrderByID(context, updateOrderRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccess(http.StatusOK, "order successfully updated", responseUpdateOrder))
}

func (c *OrderController) DeleteOrderByID(ctx *gin.Context) {
	orderDeleteRequest := new(request.DeleteOrder)
	dboError := new(helpers.Error)

	orderDeleteRequest.ID = c.ParserHelper.GetIntParam(ctx, "id")

	if orderDeleteRequest.ID == 0 {
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	dboError = c.OrderUsecase.DeleteOrderByID(context, orderDeleteRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccess(http.StatusOK, "order successfully deleted", nil))
}
