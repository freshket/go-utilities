
### Example usage

```go
err := errors.New("unauthorized")

// simple usecases
lxg.Debug("test debug log with no parameter")
lxg.Debug("test debug log", lxg.Param("p", 1))
lxg.Info("test info log", lxg.Param("p", 1))
lxg.Warn("test warn log", lxg.Param("p", 1))
lxg.Error("test error log", err, lxg.Param("p", 1))
lxg.Error2("the input is invalid", lxg.Param("p", 1))

// with struct/map param
ec2 := struct{ IP string }{
    IP: "100.1.1.1",
}
permissions := map[string]bool{
    "read":  true,
    "write": false,
}

lxg.Debug("struct param", lxg.ParamJson("ec2", ec2))
lxg.Debug("map param", lxg.ParamJson("permission", permissions))

// with multiple params
lxg.Debug("multiple params", lxg.Param("p1", 1), lxg.Param("p2", 2))

// with pointer params
s := "hi mom"
lxg.Debug("pointer param", lxg.Param("p", &s))
lxg.Debug("pointer (struct) param", lxg.ParamJson("p", &ec2))

// with correlation id
ctx := context.Background()

ctx = lxg.CtxWithCorrelationID(ctx, uuid.New().String())
lxg.DebugCtx(ctx, "with correlation id", lxg.Param("p", 1))

// with label
ctx = lxg.CtxWithLabels(
    ctx,
    lxg.Label("guid", "8b398783-d157-44d9-b850-c3b693af2a6c"),
    lxg.Label("running_number", 10),
)
ctx = lxg.CtxWithLabels(ctx, lxg.Label("restaurant_id", "666666"))

lxg.DebugCtx(ctx, "with labels", lxg.Param("test", 1))
```

```json
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "test debug log with no parameter",
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "test debug log",
  "param.p": 1,
  "ecs.version": "1.6.0"
}
{
  "log.level": "info",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "test info log",
  "param.p": 1,
  "ecs.version": "1.6.0"
}
{
  "log.level": "warn",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "test warn log",
  "param.p": 1,
  "ecs.version": "1.6.0"
}
{
  "log.level": "error",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "test error log",
  "param.p": 1,
  "error": {
    "message": "unauthorized"
  },
  "ecs.version": "1.6.0"
}
{
  "log.level": "error",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "the input is invalid",
  "param.p": 1,
  "error": {
    "message": "the input is invalid"
  },
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "struct param",
  "param.ec2": "{\"IP\":\"100.1.1.1\"}",
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "map param",
  "param.permission": "{\"read\":true,\"write\":false}",
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "multiple params",
  "param.p1": 1,
  "param.p2": 2,
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "pointer param",
  "param.p": "hi mom",
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "pointer (struct) param",
  "param.p": "{\"IP\":\"100.1.1.1\"}",
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "with correlation id",
  "correlation_id": "499e8c21-186a-4ae3-9924-0b299ec3503d",
  "param.p": 1,
  "ecs.version": "1.6.0"
}
{
  "log.level": "debug",
  "@timestamp": "2023-02-02T12:58:33.689+0700",
  "message": "with labels",
  "label.guid": "8b398783-d157-44d9-b850-c3b693af2a6c",
  "label.running_number": 10,
  "label.restaurant_id": "666666",
  "correlation_id": "499e8c21-186a-4ae3-9924-0b299ec3503d",
  "param.test": 1,
  "ecs.version": "1.6.0"
}
```