package controllers

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/helpers"
	"dbo-be-task/internal/usecases/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	UserUsecase user.UserUseCase
	ErrorHelper *helpers.ErrorHelper
	Log         *logrus.Logger
}

func NewUserController(userUsecase user.UserUseCase, errorHelper *helpers.ErrorHelper, log *logrus.Logger) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
		ErrorHelper: errorHelper,
		Log:         log,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var (
		createUserRequest *request.CreateUser
		dboError          *helpers.Error
	)

	err := ctx.ShouldBind(&createUserRequest)

	if err != nil {
		dboError = c.ErrorHelper.NewError(http.StatusBadRequest, "error while parsing JSON body request to `createUserRequest` struct", err.Error())
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseCreateUser, dboError := c.UserUsecase.CreateUser(context, createUserRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusCreated, response.NewHTTPResponseSuccess(http.StatusCreated, "user successfully registered", responseCreateUser))
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var (
		loginUserRequest *request.LoginUser
		dboError         *helpers.Error
	)
	err := ctx.ShouldBind(&loginUserRequest)

	if err != nil {
		dboError = c.ErrorHelper.NewError(http.StatusBadRequest, "error while parsing JSON body request to `loginUserRequest` struct", err.Error())
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	context, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	responseLoginUser, dboError := c.UserUsecase.LoginUser(context, loginUserRequest)

	if dboError != nil {
		ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	ctx.JSON(http.StatusOK, response.NewHTTPResponseSuccess(http.StatusOK, "user successfully login", responseLoginUser))
}
