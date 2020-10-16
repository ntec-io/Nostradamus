package logger

import "go.uber.org/zap"

var log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()
}

func Log() *zap.SugaredLogger {
	return log
}
