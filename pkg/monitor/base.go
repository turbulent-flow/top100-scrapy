package monitor

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/LiamYabou/top100-pkg/logger"
)

func Initialize() {
	if variable.Env == "development" {
		return
	}
	_, err := newrelic.NewApplication(
		newrelic.ConfigAppName(variable.AppName),
		newrelic.ConfigLicense(variable.NewRelicLicenseKey),
	)
	if err != nil {
        logger.Error("unable to create New Relic Application", err)
    }
}
