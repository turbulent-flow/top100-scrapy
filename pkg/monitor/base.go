package monitor

import (
	"github.com/newrelic/go-agent"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
)

func Initialize() {
	if variable.Env == "development" {
		return
	}
	newrelicMonitor, err := newrelic.NewApplication(
		newrelic.ConfigAppName(variable.AppName),
		newrelic.ConfigLicense(variable.NewRelicLicenseKey),
	)
	http.HandleFunc(newrelic.WrapHandleFunc(newrelicMonitor, "/users", usersHandler))
}