package zlog

import (
	"context"
	"net"
	"os"
	"strings"

	"go.opentelemetry.io/otel/trace"
)

const LOG_PH = "-" // 占位符

// customFields 检查ctx中是否存在logId，如果存在则返回，否则返回"-"
func customFields(ctx context.Context) string {
	var cStr string
	localIP := ip()
	if localIP == "" {
		localIP = LOG_PH
	}
	cStr += localIP + _defaultConsoleSeparator
	psm := env("PSM")
	if psm == "" {
		psm = LOG_PH
	}
	cStr += psm + _defaultConsoleSeparator
	e := env("ENV")
	if e == "" {
		e = LOG_PH
	}
	cStr += e + _defaultConsoleSeparator
	lid := GetTraceIDFromSpan(ctx)
	if lid == "" {
		lid = LOG_PH
	}
	cStr += lid
	return cStr
}

// 读取环境变量
func env(key string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	if os.Getenv(strings.ToUpper(key)) != "" {
		return os.Getenv(strings.ToUpper(key))
	}
	return os.Getenv(strings.ToLower(key))
}

func ip() string {
	// get local ip
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetLogIdFromCtx 对外提供: 从ctx中提取logID
func GetTraceIDFromSpan(ctx context.Context) string {
	// check param
	if ctx == nil {
		return ""
	}
	// 从请求的 context 中获取当前 span
	span := trace.SpanFromContext(ctx)
	// 获取 TraceID
	traceID := span.SpanContext().TraceID().String()
	return traceID
}
