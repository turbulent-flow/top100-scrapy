package monitor

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
)

func InitNewRelic() (newRelicApp *newrelic.Application, err error) {
	if variable.Env == "development" {
		return
	}
	newRelicApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName(variable.AppName),
		newrelic.ConfigLicense(variable.NewRelicLicenseKey),
	)
	return
}
