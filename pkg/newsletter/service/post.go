package service

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (s *service) Post(
	newElement *newsletter.Subscription,
) error {
	return s.repo.Save(newElement)
}
