package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/repository"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// nolint:lll // godoc
// Get godoc
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /subscriptions [get]
// @Param        page	        query  int		 true   "Result page"                                    example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"                                  example(10)
// @Param        userId	        query  string	 false  "User ID"                                        example(1)
// @Param        blogId	        query  string	 false  "Blog ID"                                        example(1)
// @Param        interests	    query  []string  false  "Interests"                                      example(["tech","sports"])
// @Produce      json
// @Success      200  {array}  handler.ResponseDoc
// nolint:gocyclo //error checking branches

func (h *handler) Get(ctx *gin.Context) {
	var newsLetterService newsletter.Service
	var subscriptions *newsletter.Result[*newsletter.Subscription]
	var filter = new(Filter)
	var pagination = new(Pagination)
	var err error
	var result Response
	var userID uuid.UUID
	var blogID uuid.UUID

	if pagination, err = pagination.Get(ctx); err != nil {
		result.SetErrorResponse(ctx, http.StatusBadRequest, err, "")
		return
	}

	if filter, err = filter.Get(ctx); err != nil {
		result.SetErrorResponse(ctx, http.StatusBadRequest, err, "")
		return
	}

	if filter.UserID != "" {
		if userID, err = uuid.Parse(filter.UserID); err != nil {
			result.SetErrorResponse(ctx, http.StatusBadRequest, err, "")
			return
		}
	}

	if filter.BlogID != "" {
		if blogID, err = uuid.Parse(filter.BlogID); err != nil {
			result.SetErrorResponse(ctx, http.StatusBadRequest, err, "")
			return
		}
	}

	newsLetterService = service.Must(repository.Must())

	if subscriptions, err = newsLetterService.Get(ctx, userID, blogID, filter.Interest); err != nil {
		result.SetErrorResponse(ctx, http.StatusInternalServerError, err, "")
		return
	}
	result = result.SetOkResponse(subscriptions, *filter, *pagination)
	ctx.JSON(http.StatusOK, result)
}
