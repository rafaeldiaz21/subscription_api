package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

type Response struct {
	FilterUserID    string                                       `json:"userId,omitempty"`
	FilterBlogID    string                                       `json:"blogId,omitempty"`
	FilterInterests []string                                     `json:"interests,omitempty"`
	Pagination      *Pagination                                  `json:"pagination"`
	Results         *newsletter.Result[*newsletter.Subscription] `json:"results"`
}

type Result struct {
	UserID    string   `json:"userId"`
	BlogID    string   `json:"blogId"`
	Interests []string `json:"interests"`
}

func (r Response) SetOkResponse(result *newsletter.Result[*newsletter.Subscription], filter Filter, pagination Pagination) Response {
	var response = Response{
		FilterUserID: filter.UserID,
		FilterBlogID: filter.BlogID,
		Pagination:   pagination.New(result.Total),
		Results:      result,
	}
	return response
}
