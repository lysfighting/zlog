package zlog

import (
	"context"
	"strings"

	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
)

var (
	globalLogger *ZLogger
	defaultConf  = LogConfig{
		LogPath:    "app.log",
		LogLevel:   "info",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}
)

func init() {
	globalLogger = NewZLoggerFromConfig(defaultConf, sdklog.NewLoggerProvider())
}

type ZLogger struct {
	Slogger *zap.SugaredLogger // 为了兼容其他接口所以暴露该变量，不直接操作该变量
}

func NewZLogger(_logger *zap.SugaredLogger) *ZLogger {
	_logger.Desugar()
	logger := &ZLogger{Slogger: _logger}
	globalLogger = logger
	return logger
}

func NewZLoggerFromConfig(cfg LogConfig, loggerProvider *sdklog.LoggerProvider) *ZLogger {
	logger := &ZLogger{Slogger: createZapLogger(cfg, loggerProvider).Sugar()}
	globalLogger = logger
	return logger
}

func GetLogger() *ZLogger {
	return globalLogger
}

func (l ZLogger) Info(ctx context.Context, args ...interface{}) {
	logIdArgs := []interface{}{customFields(ctx), _defaultConsoleSeparator}
	args = append(logIdArgs, args...)
	l.Slogger.Info(args...)
}

// Infof logs an info msg with fields
func (l ZLogger) Infof(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{customFields(ctx), template}, _defaultConsoleSeparator)
	l.Slogger.Infof(template, args...)
}

func (l ZLogger) Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{customFields(ctx), msg}, " ")
	l.Slogger.Infow(msg, keysAndValues...)
}

func (l ZLogger) Debug(ctx context.Context, args ...interface{}) {
	logIdArgs := []interface{}{customFields(ctx), _defaultConsoleSeparator}
	args = append(logIdArgs, args...)
	l.Slogger.Debug(args...)
}

// Debugf logs an debug msg with fields
func (l ZLogger) Debugf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{customFields(ctx), template}, _defaultConsoleSeparator)
	l.Slogger.Debugf(template, args...)
}

func (l ZLogger) Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{customFields(ctx), msg}, " ")
	l.Slogger.Debugw(msg, keysAndValues...)
}

func (l ZLogger) Error(ctx context.Context, args ...interface{}) {
	logIdArgs := []interface{}{customFields(ctx), _defaultConsoleSeparator}
	args = append(logIdArgs, args...)
	l.Slogger.Error(args...)
}

// Errorf logs an error msg with fields
func (l ZLogger) Errorf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{customFields(ctx), template}, _defaultConsoleSeparator)
	l.Slogger.Errorf(template, args...)
}

// Errorw print the err msg by json
func (l ZLogger) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{customFields(ctx), msg}, " ")
	l.Slogger.Errorw(msg, keysAndValues...)
}

func (l ZLogger) Fatal(ctx context.Context, args ...interface{}) {
	logIdArgs := []interface{}{customFields(ctx), _defaultConsoleSeparator}
	args = append(logIdArgs, args...)
	l.Slogger.Fatal(args...)
}

// Fatalf logs a fatal error msg with fields
func (l ZLogger) Fatalf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{customFields(ctx), template}, _defaultConsoleSeparator)
	l.Slogger.Fatalf(template, args...)
}

func (l ZLogger) Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{customFields(ctx), msg}, " ")
	l.Slogger.Fatalw(msg, keysAndValues...)
}

func (l ZLogger) Warn(ctx context.Context, args ...interface{}) {
	logIdArgs := []interface{}{customFields(ctx), _defaultConsoleSeparator}
	args = append(logIdArgs, args...)
	l.Slogger.Warn(args...)
}

func (l ZLogger) Warnf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{customFields(ctx), template}, _defaultConsoleSeparator)
	l.Slogger.Warnf(template, args...)
}

func (l ZLogger) Warnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{customFields(ctx), msg}, " ")
	l.Slogger.Warnw(msg, keysAndValues...)
}

// With creates a child ZLogger, and optionally adds some context fields to that ZLogger.
func (l ZLogger) With(ctx context.Context, fields ...interface{}) ZLogger {
	return ZLogger{Slogger: l.Slogger.With(fields...)}
}
