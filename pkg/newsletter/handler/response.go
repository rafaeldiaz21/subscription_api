package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
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

func (r Response) SetOkResponse(result *newsletter.Result[*newsletter.Subscription], filter Filter, pagination Pagination) Response {
	var response = Response{
		Filter:     filter,
		Pagination: pagination.New(result.Total),
		Results:    result,
	}
	return response
}
