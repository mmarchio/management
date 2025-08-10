package types

import (
	"context"

	"github.com/mmarchio/management/logger"
)

func GetLogger(ctx context.Context) logger.LoggerFuncT {
	if l, ok := ctx.Value(logger.LoggerKey).(logger.LoggerFuncT); ok {
		return l
	}
	return nil
}