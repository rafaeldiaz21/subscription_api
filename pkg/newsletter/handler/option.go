package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

type Option func(*handler) error

func WithService(svc newsletter.Service) Option {
	return func(h *handler) error {
		h.svc = svc
		return nil
	}
}
