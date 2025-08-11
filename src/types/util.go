package types

import (
	"github.com/mmarchio/management/logger"
)

func GetLogger() logger.LoggingContext {
	r := logger.LoggingContext{}
	r.Init()
	return r
}