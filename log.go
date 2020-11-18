package log

import "context"

// LoggerFactory 日志工厂的抽象
type LoggerFactory interface {
	DefaultLogger(level LogLevel) Logger
}

// Logger 日志实例的抽象
type Logger interface {
	Fatalf(ctx context.Context, format string, v ...interface{})
	Fatal(ctx context.Context, msg string)

	Errorf(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, msg string)

	Warnf(ctx context.Context, format string, v ...interface{})
	Warn(ctx context.Context, msg string)

	Infof(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, msg string)

	Debugf(ctx context.Context, format string, v ...interface{})
	Debug(ctx context.Context, msg string)

	Tracef(ctx context.Context, format string, v ...interface{})
	Trace(ctx context.Context, msg string)

	LogLevel() LogLevel
}

// WithTraceId 向context中注入traceid
func WithTraceId(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, TraceId, v)
}

func GetTraceId(ctx context.Context) string {
	if v, ok := ctx.Value(TraceId).(string); ok {
		return v
	}
	return ""
}

// WithServiceName 向context中注入servicename
func WithServiceName(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, ServiceName, v)
}

func getServiceName(ctx context.Context) string {
	if v, ok := ctx.Value(ServiceName).(string); ok {
		return v
	}
	return ""
}

// WithModuleName 向context中注入模块名
func WithModuleName(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, ModuleName, v)
}

func getModuleName(ctx context.Context) string {
	if v, ok := ctx.Value(ModuleName).(string); ok {
		return v
	}
	return ""
}

// CopyContextValue 将log上下文数据从一个context拷贝至另一个context
func CopyContextValue(from, to context.Context) (newContext context.Context) {
	newContext = to
	if v := GetTraceId(from); v != "" {
		newContext = WithTraceId(newContext, v)
	}
	if v := getServiceName(from); v != "" {
		newContext = WithServiceName(newContext, v)
	}
	if v := getModuleName(from); v != "" {
		newContext = WithModuleName(newContext, v)
	}
	return newContext
}
