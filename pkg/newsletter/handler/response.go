package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Filter     Filter                                       `json:"filter,omitempty"`
	Pagination *Pagination                                  `json:"pagination"`
	Results    *newsletter.Result[*newsletter.Subscription] `json:"results"`
}

type Result struct {
	UserID    string   `json:"userId"`
	BlogID    string   `json:"blogId"`
	Interests []string `json:"interests"`
}

type ErrorResponse struct {
	Err     string `json:"error"`
	Message string `json:"message"`
}

func (r Response) SetOkResponse(result *newsletter.Result[*newsletter.Subscription], filter Filter, pagination Pagination) Response {
	var response = Response{
		Filter:     filter,
		Pagination: pagination.New(result.Total),
		Results:    result,
	}
	return response
}

// TO DO it must be according API agreements
func (r Response) SetErrorResponse(ctx *gin.Context, code int, err error, message string) {
	var response = ErrorResponse{
		Err:     err.Error(),
		Message: message,
	}
	ctx.JSON(code, response)
}
