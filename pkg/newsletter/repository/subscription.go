package repository

import (
	newsletter "git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

type subscriptionDBModel struct {
	UserID    string
	BlogID    string
	Interests []string
}

func subscriptionDBModelToSubscription(data []*subscriptionDBModel) []*newsletter.Subscription {

	var response = make([]*newsletter.Subscription, len(data))

	for i := range data {

		var interests = make([]newsletter.Interest, len(data[i].Interests))
		for i := range data[i].Interests {
			interests[i] = newsletter.Interest(data[i].Interests[i])
		}

		response[i] = &newsletter.Subscription{
			UserID:    uuid.MustParse(data[i].UserID),
			BlogID:    uuid.MustParse(data[i].BlogID),
			Interests: interests,
		}
	}

	return response
}

func subscriptionToSubscriptionDBModel(data *newsletter.Subscription) *subscriptionDBModel {

	var interests = make([]string, len(data.Interests))
	for i := range data.Interests {
		interests[i] = string(data.Interests[i])
	}

	var response = &subscriptionDBModel{
		UserID:    data.UserID.String(),
		BlogID:    data.BlogID.String(),
		Interests: interests,
	}
	return response
}
