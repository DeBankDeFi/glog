package glog

type logFactory struct{}

func NewLogFactory() *logFactory {
	return &logFactory{}
}

func (f *logFactory) DefaultLogger(level LogLevel) Logger {
	return newLoggerV1(level)
}
