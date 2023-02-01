package repository

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (r *repository) Save(
	newElement *newsletter.Subscription,
) error {
	r.data = append(r.data, subscriptionToSubscriptionDBModel(newElement))
	return nil
}
