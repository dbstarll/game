package logger

import (
	"go.uber.org/zap"
	"log"
)

func init() {
	if logger, err := zap.NewDevelopment(zap.AddStacktrace(zap.FatalLevel)); err != nil {
		log.Fatalf("%+v", err)
	} else {
		zap.RedirectStdLog(logger)
		zap.ReplaceGlobals(logger)
	}
}
