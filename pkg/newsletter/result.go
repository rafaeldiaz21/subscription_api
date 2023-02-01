package newsletter

import (
	"context"
	"math"
)

type Result[T any] struct {
	Total int
	Pages int
	Page  Page[T]
}

func (r *Result[T]) Get(ctx context.Context, data []T) {
	var pageRequest = ctx.Value("page").(int)
	var maxPageSize = ctx.Value("maxPageSize").(int)
	var offset = (maxPageSize * (pageRequest - 1))
	var limit = maxPageSize
	var page = new(Page[T])

	r.Pages = CountPages(maxPageSize, len(data))
	r.Total = len(data)
	r.Page = page.New(data, limit, offset)
}

func CountPages(maxPageSize int, totalElements int) int {
	if totalElements == 0 {
		return 0
	}

	if totalElements <= maxPageSize {
		return 1
	}
	return int(
		math.Round(
			float64(totalElements) / float64(maxPageSize),
		),
	)
}
