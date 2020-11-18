package log

// LogLevel 日志等级
type LogLevel uint8

const (
	LevelFatal = LogLevel(iota + 1)
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)

// ContextKey 注入context中的额外信息
type ContextKey string

const (
	// TraceId 用于上下文trace的唯一ID
	TraceId = ContextKey("TraceId")

	// ServiceName 产出日志的服务名
	ServiceName = ContextKey("ServiceName")

	// ModuleName 产出日志的模块名
	ModuleName = ContextKey("ModuleName")
)
