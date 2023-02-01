package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/utils"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) (_ []*newsletter.Subscription, _ error) {
	var data []*newsletter.Subscription
	data = subscriptionDBModelToSubscription(r.data)
	data = search(data, userID, blogID)
	return data, nil
}

func search(data []*newsletter.Subscription, userID uuid.UUID, blogID uuid.UUID) []*newsletter.Subscription {
	if userID.String() == uuid.Nil.String() && blogID.String() == uuid.Nil.String() {
		return data
	}

	data = utils.Filter(data, func(el *newsletter.Subscription) bool {

		if userID.String() != uuid.Nil.String() && blogID.String() != uuid.Nil.String() {
			return userID.String() == el.UserID.String() && blogID.String() == el.BlogID.String()
		}

		if userID.String() != uuid.Nil.String() && userID.String() == el.UserID.String() {
			return true
		}
		if blogID.String() != uuid.Nil.String() && blogID.String() == el.BlogID.String() {
			return true
		}
		return false
	})

	return data
}
