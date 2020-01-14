package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"music-pc-server/internal/app/config"
)

var DBCNormalLogger *zap.Logger

func InitLog() {

	cfg := config.Global()

	var level zapcore.Level

	switch cfg.Log.Level {
	case "dev":
		level = zapcore.DebugLevel
	case "test":
		level = zapcore.DebugLevel
	case "prod":
		level = zapcore.InfoLevel
	default:
		level = zapcore.InfoLevel
	}
	DBCNormalLogger = NewLogger(cfg.Log.OutFile, level, 20, 3, 7, true)
}
