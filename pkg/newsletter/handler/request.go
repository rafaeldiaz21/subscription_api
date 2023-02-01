package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Pagination *Pagination
	Filter     *Filter
}

func Get(ctx *gin.Context) (*request, error) {
	var fill = new(request)
	var err error
	if err = ctx.BindQuery(fill); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return nil, err
	}

	ctx.Set("page", fill.Pagination.Page)
	ctx.Set("maxPageSize", fill.Pagination.MaxPageSize)
	return fill, nil
}
