# CLAUDE.md

本文件为 Claude Code（claude.ai/code）在此仓库中工作时提供指导。

## 构建与运行

```bash
go build -o server ./cmd/   # 构建（包含 wire_gen.go）
go run ./cmd/               # 运行（需要 MySQL）
```

项目中没有 Makefile，也没有测试文件——后续添加测试时，直接在对应包目录下创建 `foo_test.go` 即可。

## 架构

这是一个早期的 Go Web API 脚手架，使用 **Gin** + **GORM** + **Viper** + **google/wire**，采用分层架构：

```
handler (HTTP) → service (业务逻辑) → repository (数据访问)
```

- **`cmd/main.go`** — 入口：调用 `InitializeApp()` 获取依赖，auto-migrate，启动服务。
- **`cmd/wire.go`** — wire 注入器声明（`//go:build wireinject`），定义 provider 和 App struct。
- **`cmd/wire_gen.go`** — wire 生成的组装代码（`//go:build !wireinject`），**不要手动编辑**。
- **`config/`** — 基于 Viper 的 YAML 配置加载。新增 `JWTConfig` 需单独注入。
- **`internal/server/`** — 创建 Gin 引擎并注册路由。
- **`internal/handler/`** — 按领域拆文件。`handler.go` 为公共 `Handler`（Ping），`user.go` 为 `UserHandler`（Register/Login）。`base.go` 提供 `Success/Error` JSON 响应方法。
- **`internal/service/`** — 按领域拆文件。`service.go` 为公共 `Service`，`user.go` 为 `UserService`。
- **`internal/repository/`** — 按领域拆文件。`repository.go` 为公共 `Repository`，`user.go` 为 `UserRepo`。
- **`internal/router/`** — `router.go` 创建 engine 和公共路由，按领域拆路由注册函数（如 `user.go` 的 `registerUserRoutes`）。
- **`internal/model/`** — GORM 模型（按数据库表结构生成）。

### 依赖注入（google/wire）

wire 在编译时生成依赖注入代码——修改任何构造函数签名后需重新运行：

```bash
wire ./cmd/...
```

每条依赖链路独立：`Handler → Service → Repository` 和 `UserHandler → UserService → UserRepo` 互不干扰。

### 新增领域模块的步骤

1. `internal/model/` — 新增 GORM 模型
2. `internal/repository/foo.go` — `FooRepo` struct + 数据访问方法
3. `internal/service/foo.go` — `FooService` struct + 业务逻辑
4. `internal/handler/foo.go` — `FooHandler` struct + HTTP handler
5. `internal/router/foo.go` — `registerFooRoutes(rg, h)`
6. `cmd/wire.go` — 添加新 provider 到 `wire.Build`
7. `cmd/main.go` — 将 `FooHandler` 传给 `server.Start`
8. 运行 `wire ./cmd/...` 重新生成

## 配置

启动时读取 `config/config.yaml`。支持环境变量覆盖（前缀 `BI_`，如 `BI_JWT_SECRET=xxx`）。

## Docker

```bash
docker build -t bi-api .
docker run -p 8080:8080 bi-api
```

配置文件中的 MySQL 地址为 `127.0.0.1:3306`，在容器内部无法解析到宿主机——本地 Docker 开发时需改用 `host.docker.internal` 或编写带 MySQL 服务的 compose 文件。
