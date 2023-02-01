package handler

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page             int    `json:"page" form:"page" binding:"required"`
	MaxPageSize      int    `json:"maxPageSize" form:"maxPageSize" binding:"required"`
	NumberOfPages    int    `json:"numberOfPages"`
	PaginationString string `json:"paginationString"`
	TotalElements    int    `json:"totalElements"`
}

func (p *Pagination) Get(ctx *gin.Context) (pagination *Pagination, err error) {
	var pag = new(Pagination)

	if err = ctx.BindQuery(pag); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	ctx.Set("page", pag.Page)
	ctx.Set("maxPageSize", pag.MaxPageSize)

	return pag, nil
}

func (p *Pagination) New(totalElements int) *Pagination {
	var numberOfPages = CountPages(p.MaxPageSize, totalElements)
	return &Pagination{
		Page:             p.Page,
		PaginationString: fmt.Sprintf("%v/%v", p.Page, numberOfPages),
		MaxPageSize:      p.MaxPageSize,
		NumberOfPages:    CountPages(p.MaxPageSize, totalElements),
		TotalElements:    totalElements,
	}
}

func CountPages(maxPageSize int, totalElements int) int {
	if totalElements == 0 {
		return 0
	}

	if totalElements <= maxPageSize {
		return 1
	}
	return int(math.Round(float64(totalElements) / float64(maxPageSize)))
}
