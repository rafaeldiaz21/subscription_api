package service

import (
	"sync"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

var (
	srvc newsletter.Service
	once sync.Once
)

type service struct {
	repo newsletter.Repository
}

func Must(
	repo newsletter.Repository,
) newsletter.Service {
	once.Do(func() {
		srvc = &service{
			repo: repo,
		}
	})

	return srvc
}
