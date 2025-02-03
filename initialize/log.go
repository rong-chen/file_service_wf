package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogContext() (log *zap.Logger) {
	logFile, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	writeSyncer2 := zapcore.AddSync(logFile)
	async := zapcore.NewMultiWriteSyncer(writeSyncer2, zapcore.AddSync(os.Stdout))
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, async, zapcore.InfoLevel)
	return zap.New(core)
}
