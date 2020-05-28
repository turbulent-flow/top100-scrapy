package logger

// It's a new wrapper of logger, extends several general methods.

import (
	log "github.com/sirupsen/logrus"
	"fmt"
)

// Construct the factors as the first item of args passed into the methods.
type Factors map[string]interface{}

// args: Receive the factors as the first item of args, it's optional.
// e.g.
// resp, err := http.Get(url)
// if err != nil {
// 	logger.Error("Failed to get the url.", err)
// }
// defer resp.Body.Close()
//
// if resp.StatusCode != 200 {
// 	factors := logger.Factors{"status_code": resp.StatusCode, "status": resp.Status}
// 	logger.Error("The status of the code error occurs!", err, factors)
// }
func Error(msg string, err error, args ...Factors) {
	msg = fmt.Sprintf("%s Error: %s", msg, err)
	switchError(NewFactorsEntry(args...), msg)
}

func Info(msg string, args ...Factors) {
	NewFactorsEntry(args...).Info(msg)
}

func NewFactorsEntry(args ...Factors) *log.Entry {
	l := log.Fields{}
	if len(args) > 0 {
		f := args[0]
		for k, v := range f {
			l[k] = v
		}
	}
	return log.WithFields(l)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}
