package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitLogger() {
	zapLog, _ := zap.NewDevelopment()
	Log = zapLog.Sugar()
}
