package model

import (
	"github.com/LiamYabou/top100-pkg/logger"
)

type OurOfIndexError struct {
	error
	Factors logger.Factors
}
