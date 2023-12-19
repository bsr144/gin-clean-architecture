package presenters

import (
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/usecases/user"

	"github.com/sirupsen/logrus"
)

type userPresenter struct {
	Log *logrus.Logger
}

func NewUserPresenter(log *logrus.Logger) user.UserPresenter {
	return &userPresenter{
		Log: log,
	}
}

func (p *userPresenter) PresentLoginUser(accessToken string) *response.LoginUser {
	return &response.LoginUser{
		AccessToken: accessToken,
	}
}

func (p *userPresenter) PresentCreateUser(entityUser *entities.User) *response.CreateUser {
	return &response.CreateUser{
		ID:    entityUser.ID,
		Email: entityUser.Email,
	}
}
