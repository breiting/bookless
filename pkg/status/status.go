package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status structure with code and message presentable to the user
type Status struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewStatus creates a new object by the given information
func NewStatus(code int, message string) *Status {
	status := &Status{
		Code:    code,
		Message: message,
	}
	return status
}

// NewHTTPStatus encapsulates a proper http error response
func NewHTTPStatus(ctx *gin.Context, status int, err error) {
	er := Status{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// Send sends the status back as a JSON response
func (s *Status) Send(ctx *gin.Context) {
	ctx.JSON(s.Code, s)
}

// Implements the error interface
func (s Status) Error() string {
	if s.Message != "" {
		return s.Message
	}
	return http.StatusText(s.Code)
}
