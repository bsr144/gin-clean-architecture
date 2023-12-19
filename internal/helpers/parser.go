package helpers

import (
	"dbo-be-task/internal/adapters/dto/request"
	"dbo-be-task/internal/adapters/dto/response"
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

func (parser *ParserHelper) BindJSON(ctx *gin.Context, obj interface{}) *Error {
	if err := ctx.ShouldBind(&obj); err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusBadRequest, "error while parsing JSON body", err.Error())
		parser.handleError(ctx, dboError)
		return dboError
	}
	return nil
}

func (parser *ParserHelper) GetIntParam(ctx *gin.Context, paramName string) int {
	val, err := strconv.Atoi(ctx.Param(paramName))
	if err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusBadRequest, "error parsing parameter to integer", err.Error())
		parser.handleError(ctx, dboError)
		return 0
	}
	return val
}

func (parser *ParserHelper) GetIntCtx(ctx *gin.Context, key string) int {
	val, err := strconv.Atoi(ctx.GetString(key))
	if err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusBadRequest, "error parsing context value to integer", err.Error())
		parser.handleError(ctx, dboError)
		return 0
	}
	return val
}

func (parser *ParserHelper) BindQueryParams(ctx *gin.Context, queries *request.Query) *Error {
	err := ctx.ShouldBindQuery(&queries)
	if err != nil {
		dboError := parser.ErrorHelper.NewError(http.StatusBadRequest, "error parsing query parameters", err.Error())
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
	ctx.JSON(dboError.Code, response.NewHTTPResponseError(dboError.Code, dboError))
}
