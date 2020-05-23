package crawler

import (
	"top100-scrapy/pkg/logger"
)

type EmptyError struct {
	error
	Factors logger.Factors
}
