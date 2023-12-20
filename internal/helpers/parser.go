package helpers

import (
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ParserHelper struct {
	ErrorHelper *ErrorHelper
}

func NewParserHelper(errorHelper *ErrorHelper) *ParserHelper {
	return &ParserHelper{
		ErrorHelper: errorHelper,
	}
}

// Parse incoming JSON body request to a struct,
// 'structName' argument used for error handling information.
func (parser *ParserHelper) BindJSON(ctx *gin.Context, objectDestination interface{}, structName string) *Error {
	if err := ctx.ShouldBind(&objectDestination); err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusUnprocessableEntity, fmt.Sprintf("error while parsing JSON body to '%s' struct", structName), err.Error())
		parser.handleError(ctx, dboError)
		return dboError
	}
	return nil
}

// Get integer value of incoming param path variable,
// Example: /users/:id - paramName is 'id' and we want to change it into integer.
func (parser *ParserHelper) GetIntParam(ctx *gin.Context, paramName string) int {
	val, err := strconv.Atoi(ctx.Param(paramName))
	if err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusUnprocessableEntity, fmt.Sprintf("error while parsing parameter '%s' to integer", paramName), err.Error())
		parser.handleError(ctx, dboError)
		return 0
	}
	return val
}

// Get integer value from key-value that previously added via gin `ctx.Set`
func (parser *ParserHelper) GetIntCtx(ctx *gin.Context, key string) int {
	val, err := strconv.Atoi(ctx.GetString(key))
	if err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusUnprocessableEntity, fmt.Sprintf("error while parsing context value '%s' to integer", key), err.Error())
		parser.handleError(ctx, dboError)
		return 0
	}
	return val
}

// Parse incoming query params to a query struct,
// Example: /users?search=test&page=1&size=5 - search, page, and size are the query params we want to parse.
func (parser *ParserHelper) BindQueryParams(ctx *gin.Context, queries *request.Query) *Error {
	err := ctx.ShouldBindQuery(&queries)
	if err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusUnprocessableEntity, "error while parsing query parameters", err.Error())
		parser.handleError(ctx, dboError)
		return dboError
	}

	// Set default values
	if queries.Page == 0 {
		queries.Page = 1
	}
	if queries.Size == 0 {
		queries.Size = 10
	}

	return nil
}

func (parser *ParserHelper) handleError(ctx *gin.Context, dboError *Error) {
	ctx.AbortWithStatusJSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
}
