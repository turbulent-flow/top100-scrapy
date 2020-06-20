package kit

import (
	"time"
	"fmt"
	"github.com/LiamYabou/top100-pkg/logger"
)

type retryCallback func() (err error)

func Retry(attemps int, sleep time.Duration, c retryCallback) (err error) {
	for i := 0; ; i++ {
		err = c()
		if err == nil {
			return
		}
		if i >= (attemps - 1) {
			break
		}
		time.Sleep(sleep)
		msg := fmt.Sprintf("Retry after error: %s", err)
		logger.Info(msg)
	}
	return fmt.Errorf("After %d attempts, last error: %s", attemps, err)
}