package helpers

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

type Error struct {
	Code     int           `json:"code"`
	Message  string        `json:"message"`
	Info     string        `json:"info"`
	Detail   string        `json:"detail"`
	Location ErrorLocation `json:"location"`
}

type ErrorLocation struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

type ErrorHelper struct {
	Log *logrus.Logger
}

func NewErrorHelper(log *logrus.Logger) *ErrorHelper {
	return &ErrorHelper{
		Log: log,
	}
}

func (e *Error) Error() string {
	return e.Detail
}

// 'Info' parameter is used for more friendly error message
// 'Detail' parameter is used for more detailed error message <usually comes from the package or function used directly, e.g err.Error()>
func (e *ErrorHelper) NewError(code int, info, detail string) *Error {
	errorCategory := "Client side error"

	if code >= statusCodeMessageServerMin {
		errorCategory = "Server side error"
	}

	e.Log.WithFields(logrus.Fields{
		"error": detail,
		"code":  code,
	}).Error(errorCategory)

	_, file, line, _ := runtime.Caller(1)

	err := &Error{
		Code:    code,
		Message: e.statusMessage(code),
		Info:    info,
		Detail:  detail,
		Location: ErrorLocation{
			File: file,
			Line: line,
		},
	}

	return err
}

// This function copied from fiber framework utils package
// Thanks to fiber utils package, could be accessed via: https://github.com/gofiber/fiber/blob/v2.51.0/utils
func (e *ErrorHelper) statusMessage(code int) string {
	if code < statusCodeMessageMin || code > statusCodeMessageMax {
		return ""
	}
	return statusCodeMessage[code]
}
