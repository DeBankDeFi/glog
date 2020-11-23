// +build unittest

package glog

import (
	"context"
	"testing"
)

func TestLoggerv1(t *testing.T) {
	ctx := prepareContext()

	newLoggerV1(LevelTrace).Tracef(ctx, "Hello %s, this is tracef", "guys")
	newLoggerV1(LevelTrace).Trace(ctx, "Hello guys, this is trace")

	newLoggerV1(LevelDebug).Debugf(ctx, "Hello %s, this is debugf", "guys")
	newLoggerV1(LevelDebug).Debug(ctx, "Hello guys, this is debug")

	newLoggerV1(LevelInfo).Infof(ctx, "Hello %s, this is infof", "guys")
	newLoggerV1(LevelInfo).Info(ctx, "Hello guys, this is info")

	newLoggerV1(LevelWarn).Warnf(ctx, "Hello %s, this is warnf", "guys")
	newLoggerV1(LevelWarn).Warn(ctx, "Hello guys, this is warn")

	newLoggerV1(LevelError).Errorf(ctx, "Hello %s, this is errorf", "guys")
	newLoggerV1(LevelError).Error(ctx, "Hello guys, this is error")

	newLoggerV1(0).Fatalf(ctx, "Hello %s, this is fatalf", "guys")
	newLoggerV1(0).Fatal(ctx, "Hello guys, this is fatal")
}

func prepareContext() context.Context {
	ctx := context.Background()
	ctx = WithTraceId(ctx, "696ce42f-a56c-40a7-accf-035671b81ca6")
	ctx = WithServiceName(ctx, "alertmanager-webhook")
	ctx = WithModuleName(ctx, "pkg.log")
	return ctx
}

func BenchmarkLoggerv1(b *testing.B) {
	b.ReportAllocs()

	ctx := prepareContext()
	logger := newLoggerV1(LevelDebug)

	for i := 0; i < b.N; i++ {
		logger.Debugf(ctx, "Hello %s, this is debugf", "guys")
	}
}
