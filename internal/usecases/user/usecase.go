package user

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/helpers"
	"net/http"

	"github.com/sirupsen/logrus"
)

type usecase struct {
	UserRepository UserRepository
	UserPresenter  UserPresenter
	SecurityHelper *helpers.SecurityHelper
	ErrorHelper    *helpers.ErrorHelper
	Log            *logrus.Logger
}

func NewUserUsecase(repository UserRepository, presenter UserPresenter, securityHelper *helpers.SecurityHelper, errorHelper *helpers.ErrorHelper, log *logrus.Logger) UserUseCase {
	return &usecase{
		UserRepository: repository,
		UserPresenter:  presenter,
		SecurityHelper: securityHelper,
		ErrorHelper:    errorHelper,
		Log:            log,
	}
}

func (u *usecase) CreateUser(ctx context.Context, createUserRequest *request.CreateUser) (*response.CreateUser, *helpers.Error) {
	var err error

	user, dboError := u.UserRepository.GetUserByEmail(ctx, createUserRequest.Email)

	if dboError != nil {
		return nil, dboError
	}

	if user.IsExist() {
		return nil, u.ErrorHelper.NewError(http.StatusBadRequest, "user already exist", "user already exist")
	}

	// Set user email & password to the requested email &password
	user.Email = createUserRequest.Email
	user.Password = createUserRequest.Password

	// Generate unique salt
	user.Salt, err = u.SecurityHelper.GenerateSalt()

	if err != nil {
		return nil, u.ErrorHelper.NewError(http.StatusInternalServerError, "error while generating user salt", err.Error())
	}

	// Generate unique password based on user.Salt, then change user.Password to the new hashed password
	user.Password = u.SecurityHelper.HashPassword(user.Password, user.Salt)

	user, dboError = u.UserRepository.CreateUser(ctx, user)

	if dboError != nil {
		return nil, dboError
	}

	return u.UserPresenter.PresentCreateUser(user), nil
}

func (u *usecase) LoginUser(ctx context.Context, loginUserRequest *request.LoginUser) (*response.LoginUser, *helpers.Error) {
	user, dboError := u.UserRepository.GetUserByEmail(ctx, loginUserRequest.Email)

	if dboError != nil {
		return nil, dboError
	}

	// Check if user doesn't exist
	if !user.IsExist() {
		return nil, u.ErrorHelper.NewError(http.StatusBadRequest, "user doesn't exist", "user doesn't exist")
	}

	// Check if given password is the same as password in database
	isPasswordValid := u.SecurityHelper.ComparePassword(user.Password, user.Salt, loginUserRequest.Password)

	if !isPasswordValid {
		return nil, u.ErrorHelper.NewError(http.StatusBadRequest, "the password given doesn't match", "the password given doesn't match")
	}

	userAccessToken, err := u.SecurityHelper.GenerateToken(user.ID)

	if err != nil {
		return nil, u.ErrorHelper.NewError(http.StatusInternalServerError, "error while generating user token", err.Error())
	}

	return u.UserPresenter.PresentLoginUser(userAccessToken), nil
}
