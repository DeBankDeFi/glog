package glog

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type loggerv1 struct {
	level  LogLevel
	logger *log.Logger
}

func newLoggerV1(level LogLevel) *loggerv1 {
	return &loggerv1{
		level:  level,
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *loggerv1) Fatalf(ctx context.Context, format string, v ...interface{}) {
	if l.level >= LevelFatal {
		l.logger.Println(l.format(ctx, "FATAL", fmt.Sprintf(format, v...)))
		os.Exit(1)
	}
}

func (l *loggerv1) Fatal(ctx context.Context, msg string) {
	if l.level >= LevelFatal {
		l.logger.Println(l.format(ctx, "FATAL", msg))
		os.Exit(1)
	}
}

func (l *loggerv1) Errorf(ctx context.Context, format string, v ...interface{}) {
	if l.level >= LevelError {
		l.logger.Println(l.format(ctx, "ERROR", fmt.Sprintf(format, v...)))
	}
}

func (l *loggerv1) Error(ctx context.Context, msg string) {
	if l.level >= LevelError {
		l.logger.Println(l.format(ctx, "ERROR", msg))
	}
}

func (l *loggerv1) Warnf(ctx context.Context, format string, v ...interface{}) {
	if l.level >= LevelWarn {
		l.logger.Println(l.format(ctx, "WARN", fmt.Sprintf(format, v...)))
	}
}

func (l *loggerv1) Warn(ctx context.Context, msg string) {
	if l.level >= LevelWarn {
		l.logger.Println(l.format(ctx, "WARN", msg))
	}
}

func (l *loggerv1) Infof(ctx context.Context, format string, v ...interface{}) {
	if l.level >= LevelInfo {
		l.logger.Println(l.format(ctx, "INFO", fmt.Sprintf(format, v...)))
	}
}

func (l *loggerv1) Info(ctx context.Context, msg string) {
	if l.level >= LevelInfo {
		l.logger.Println(l.format(ctx, "INFO", msg))
	}
}

func (l *loggerv1) Debugf(ctx context.Context, format string, v ...interface{}) {
	if l.level >= LevelDebug {
		l.logger.Println(l.format(ctx, "DEBUG", fmt.Sprintf(format, v...)))
	}
}

func (l *loggerv1) Debug(ctx context.Context, msg string) {
	if l.level >= LevelDebug {
		l.logger.Println(l.format(ctx, "DEBUG", msg))
	}
}

func (l *loggerv1) Tracef(ctx context.Context, format string, v ...interface{}) {
	if l.level >= LevelTrace {
		l.logger.Println(l.format(ctx, "TRACE", fmt.Sprintf(format, v...)))
	}
}

func (l *loggerv1) Trace(ctx context.Context, msg string) {
	if l.level >= LevelTrace {
		l.logger.Println(l.format(ctx, "TRACE", msg))
	}
}

func (l *loggerv1) LogLevel() LogLevel {
	return l.level
}

func (l *loggerv1) format(ctx context.Context, levelFlag string, output string) string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		file = filepath.Base(file)
	}

	if traceId := GetTraceId(ctx); traceId != "" {
		return fmt.Sprintf(" %-3s:%-4d %-5s --- [ %-6s | %-6s | %-6s ] :   %s", file, line, levelFlag, getServiceName(ctx), getModuleName(ctx), traceId, output)
	}
	return fmt.Sprintf(" %-3s:%-4d %-5s --- [ %-6s | %-6s ] :   %s", file, line, levelFlag, getServiceName(ctx), getModuleName(ctx), output)
}
