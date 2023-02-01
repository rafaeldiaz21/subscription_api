package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	Interests []newsletter.Interest,
) (*newsletter.Result[*newsletter.Subscription], error) {
	var page = ctx.Value("page").(int)
	var maxPageSize = ctx.Value("maxPageSize").(int)
	var offset = (maxPageSize * (page - 1))
	var limit = maxPageSize
	var newsLetterSuscription []*newsletter.Subscription
	var err error
	var result *newsletter.Result[*newsletter.Subscription] = &newsletter.Result[*newsletter.Subscription]{}

	if newsLetterSuscription, err = s.repo.Search(ctx, userID, blogID, Interests, limit, offset); err != nil {
		return nil, err
	}

	result.Get(ctx, newsLetterSuscription)
	return result, nil
}
