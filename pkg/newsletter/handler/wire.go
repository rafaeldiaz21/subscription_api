//go:build wireinject
// +build wireinject

package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/repository"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/service"
	"github.com/google/wire"
)

func Build() newsletter.Handler {
	wire.Build(
		Must,
		service.Must,
		repository.Must,
	)

	return nil
}
