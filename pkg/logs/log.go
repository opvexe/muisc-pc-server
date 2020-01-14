package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DBCNormalLogger *zap.Logger

func InitLog() {

	var level zapcore.Level

	switch CurrentEnvParam {
	case "dev":
		level = zapcore.DebugLevel
	case "test":
		level = zapcore.DebugLevel
	case "prod":
		level = zapcore.InfoLevel
	default:
		level = zapcore.InfoLevel
	}
	DBCNormalLogger = NewLogger("/Logs/service_bussiness.log", level, 20, 3, 7, true)
}
