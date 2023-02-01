package handler

import (
	"fmt"
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/repository"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/service"
	"github.com/gin-gonic/gin"
)

// Post godoc
// @Summary     Create subscriptions
// @Description Remove the suspended state from user
// @Tags        subscriptions
// @Accept      json
// @Produce     json
// @Success     200 {object} newsletter.Subscription
// @Router      /subscriptions [post]

func (h *handler) Post(ctx *gin.Context) {
	var request = new(newsletter.Subscription)
	var newsLetterService newsletter.Service
	var err error
	newsLetterService = service.Must(repository.Must())

	if err = ctx.BindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("bad request, %v", err))
		return
	}

	if err = newsLetterService.Post(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error, %v", err))
		return
	}

	ctx.JSON(http.StatusCreated, request)
}
