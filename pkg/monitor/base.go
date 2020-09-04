package monitor

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/getsentry/sentry-go"
	"os"
)

func InitNewRelic() (newRelicApp *newrelic.Application, err error) {
	if variable.Env == "development" {
		return
	}
	// newRelicLogPath := fmt.Sprintf("%s/logs/go-agent.log", variable.AppURI)
	// w, err := os.OpenFile(newRelicLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// factors := logger.Factors{
	// 	"newRelicLogPath":  newRelicLogPath,
	// }
	// logger.Info("troubleshoot the newRelic", factors)
	// if nil == err {
	// 	newrelic.NewDebugLogger(w)
	// }
	newRelicApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName(variable.AppName),
		newrelic.ConfigLicense(variable.NewRelicLicenseKey),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	return
}

func InitSentry() (err error) {
	if variable.Env == "development" {
		return
	}
	err = sentry.Init(sentry.ClientOptions{})
	return
}
