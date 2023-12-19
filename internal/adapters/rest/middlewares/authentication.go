package middlewares

import (
	"dbo-be-task/internal/adapters/dto/response"
	"dbo-be-task/internal/config"
	"dbo-be-task/internal/helpers"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	ErrorHelper *helpers.ErrorHelper
	Config      *config.SecurityConfig
	Log         *logrus.Logger
}

func NewAuthMiddleware(errorHelper *helpers.ErrorHelper, config *config.SecurityConfig, log *logrus.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		ErrorHelper: errorHelper,
		Config:      config,
		Log:         log,
	}
}

func (m *AuthMiddleware) VerifyAuth(ctx *gin.Context) {
	var dboError *helpers.Error

	tokenString := m.extractToken(ctx)

	if tokenString == "" {
		dboError = m.ErrorHelper.NewError(http.StatusUnauthorized, "authorization token not provided", "authorization token not provided")
		ctx.AbortWithStatusJSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	verifiedToken, err := m.VerifyToken(tokenString)

	if err != nil {
		dboError = m.ErrorHelper.NewError(http.StatusUnauthorized, "token signature is invalid", err.Error())

		ctx.AbortWithStatusJSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
		return
	}

	if claims, ok := verifiedToken.Claims.(jwt.MapClaims); ok && verifiedToken.Valid {
		if !m.isTokenExpired(claims) {
			userId := claims["user_id"].(float64)
			strUserId := strconv.FormatFloat(userId, 'f', -1, 64)

			ctx.Set("user_id", strUserId)
			ctx.Next()
			return
		}
	}

	dboError = m.ErrorHelper.NewError(http.StatusUnauthorized, "invalid or expired token", "invalid or expired token")
	ctx.AbortWithStatusJSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
}

func (m *AuthMiddleware) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(m.Config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, err
}

func (m *AuthMiddleware) extractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer ") {
		return strings.TrimPrefix(bearerToken, "Bearer ")
	}
	return ""
}

func (m *AuthMiddleware) isTokenExpired(claims jwt.MapClaims) bool {
	if exp, ok := claims["exp"].(float64); ok {
		return time.Unix(int64(exp), 0).Before(time.Now())
	}
	return true
}
