package zlog

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.opentelemetry.io/contrib/bridges/otelzap"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	_defaultConsoleSeparator = "\t"

	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

var zLogger *zap.SugaredLogger

func createZapLogger(cfg LogConfig, loggerProvider *sdklog.LoggerProvider) *zap.Logger {
	// log time
	encoderConfig := zap.NewProductionEncoderConfig()
	var recordTimeFormat = "2006-01-02 15:04:05.000"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.TimeKey = "created_at"
	encoderConfig.ConsoleSeparator = _defaultConsoleSeparator
	var encoder = zapcore.NewConsoleEncoder(encoderConfig)
	//var encoder = zapcore.NewJSONEncoder(encoderConfig) // 使用 json 编码器
	// log path
	logFile := cfg.LogPath
	if logFile == "" {
		logFile = "app.log" // 默认日志文件路径
	}
	fmt.Println("logFile:", logFile)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	// log level
	logLevel := zap.InfoLevel
	switch strings.ToLower(cfg.LogLevel) {
	case DebugLevel:
		logLevel = zap.DebugLevel
	case InfoLevel:
		logLevel = zap.InfoLevel
	case WarnLevel:
		logLevel = zap.WarnLevel
	case ErrorLevel:
		logLevel = zap.ErrorLevel
	}

	// output
	writers := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)}
	// new zap core
	zapCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writers...), logLevel)
	otelCore := otelzap.NewCore("zlog", otelzap.WithLoggerProvider(loggerProvider))
	teeCore := zapcore.NewTee(zapCore, otelCore)
	return zap.New(teeCore, zap.AddCaller(), zap.AddCallerSkip(1))
}
