package cmd

import (
	"github.com/go-pg/pg/v10"
	"github.com/gomodule/redigo/redis"
	"github.com/monetr/monetr/pkg/application"
	"github.com/monetr/monetr/pkg/billing"
	"github.com/monetr/monetr/pkg/config"
	"github.com/monetr/monetr/pkg/controller"
	"github.com/monetr/monetr/pkg/internal/platypus"
	"github.com/monetr/monetr/pkg/internal/stripe_helper"
	"github.com/monetr/monetr/pkg/jobs"
	"github.com/monetr/monetr/pkg/mail"
	"github.com/monetr/monetr/pkg/metrics"
	"github.com/monetr/monetr/pkg/secrets"
	"github.com/monetr/monetr/pkg/ui"
	"github.com/sirupsen/logrus"
)

func getControllers(
	log *logrus.Entry,
	configuration config.Configuration,
	db *pg.DB,
	job jobs.JobManager,
	plaidClient platypus.Platypus,
	stats *metrics.Stats,
	stripe stripe_helper.Stripe,
	cache *redis.Pool,
	plaidSecrets secrets.PlaidSecretsProvider,
	basicPaywall billing.BasicPayWall,
	smtpCommunication mail.Communication,
) []application.Controller {
	return []application.Controller{
		controller.NewController(
			log,
			configuration,
			db,
			job,
			plaidClient,
			stats,
			stripe,
			cache,
			plaidSecrets,
			basicPaywall,
			smtpCommunication,
		),
		ui.NewUIController(),
	}
}
