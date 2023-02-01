package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
)

type Filter struct {
	UserID   string                `json:"userId" form:"userId"`
	BlogID   string                `json:"blogId" form:"blogId"`
	Interest []newsletter.Interest `form:"interests"`
}

func (f *Filter) Get(ctx *gin.Context) (filter *Filter, err error) {
	var fill = new(Filter)

	if err = ctx.BindQuery(fill); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	return fill, nil
}
