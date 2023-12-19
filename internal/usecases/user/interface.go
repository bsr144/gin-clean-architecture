package user

import (
	"context"
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/helpers"
)

type UserRepository interface {
	CreateUser(ctx context.Context, entityUser *entities.User) (*entities.User, *helpers.Error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, *helpers.Error)
}

type UserPresenter interface {
	PresentCreateUser(entityUser *entities.User) *response.CreateUser
	PresentLoginUser(token string) *response.LoginUser
}

type UserUseCase interface {
	CreateUser(ctx context.Context, createUserRequest *request.CreateUser) (*response.CreateUser, *helpers.Error)
	LoginUser(ctx context.Context, loginUserRequest *request.LoginUser) (*response.LoginUser, *helpers.Error)
}
