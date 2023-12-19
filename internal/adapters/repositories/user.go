package repositories

import (
	"context"
	"database/sql"
	"dbo-be-task/internal/entities"
	"dbo-be-task/internal/helpers"
	"dbo-be-task/internal/usecases/user"
	"net/http"

	"github.com/sirupsen/logrus"
)

type userRepository struct {
	ErrorHelper *helpers.ErrorHelper
	DB          *sql.DB
	Log         *logrus.Logger
}

func NewUserRepository(errorHelper *helpers.ErrorHelper, db *sql.DB, log *logrus.Logger) user.UserRepository {
	return &userRepository{
		ErrorHelper: errorHelper,
		DB:          db,
		Log:         log,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, entityUser *entities.User) (*entities.User, *helpers.Error) {
	err := r.DB.
		QueryRowContext(ctx, CREATE_NEW_USER_WITH_RETURNING_QUERY, entityUser.Email, entityUser.Password, entityUser.Salt).
		Scan(
			&entityUser.ID,
			&entityUser.Email,
		)

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while create user on database", err.Error())
	}

	return entityUser, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, *helpers.Error) {
	entityUser := new(entities.User)

	err := r.DB.
		QueryRowContext(ctx, GET_USER_BY_EMAIL_QUERY, email).
		Scan(&entityUser.ID,
			&entityUser.Email,
			&entityUser.Password,
			&entityUser.Salt,
			&entityUser.CreatedAt,
			&entityUser.UpdatedAt,
			&entityUser.DeletedAt,
		)

	if err == sql.ErrNoRows {
		return entityUser, nil
	}

	if err != nil {
		return nil, r.ErrorHelper.NewError(http.StatusInternalServerError, "error while get user by email from database", err.Error())
	}

	return entityUser, nil
}
