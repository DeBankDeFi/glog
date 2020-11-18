# Log

A predefined Log Library for quick access when developing projects. Interface-oriented programming.


# Features

* Interface-oriented programming, easily extendable by implementation.
* The implementation is based on the packaging of Golang's built-in log library, which is simple and has no other dependencies.
* Support for injecting additional contextual parameters via Context, e.g. TraceId, service name, module name, etc., makes it easy to search when using SLS for problem location.


# Quickstart

```Go
func Demo() {
    ctx := log.WithTraceId(context.Background(), "696ce42f-a56c-40a7-accf-035671b81ca6")
    logger := log.NewLogFactory().DefaultLogger(log.LevelDebug)
    logger.Info(ctx, "hello guys, this is a info log")
}
```
