package crawler

import (
	"github.com/LiamYabou/top100-pkg/logger"
)

type EmptyError struct {
	error
	Factors logger.Factors
}
