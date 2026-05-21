# CLAUDE.md

本文件为 Claude Code（claude.ai/code）在此仓库中工作时提供指导。

## 构建与运行

```bash
go build -o server ./cmd/main.go   # 构建
go run ./cmd/main.go               # 运行（需要 MySQL）
```

项目中没有 Makefile，也没有测试文件——后续添加测试时，直接在对应包目录下创建 `foo_test.go` 即可。

## 架构

这是一个早期的 Go Web API 脚手架，使用 **Gin** + **GORM** + **Viper**，采用分层架构，在 `cmd/main.go` 中手动完成依赖注入：

```
handler (HTTP) → service (业务逻辑) → repository (数据访问)
```

- **`cmd/main.go`** — 入口：加载配置、通过 GORM 连接 MySQL、组装依赖、启动服务。
- **`config/`** — 基于 Viper 的 YAML 配置加载。`MySQLConfig.DSN()` 构建 MySQL 连接串，`setDefaults()` 填充 host/port 默认值。
- **`internal/server/`** — 创建 Gin 引擎并注册路由。目前仅有 `GET /ping`。
- **`internal/handler/`** — `Base` 提供 `Success(c, msg, data)` 和 `Error(c, code, msg)` JSON 响应方法。状态码：0=成功, 1=失败, 401, 404, 406, 500, 502。`Handler` 持有 `*service.Service` 引用。
- **`internal/service/`** — 业务逻辑层，目前为透传，后续在此编写业务代码。
- **`internal/repository/`** — 封装 `*gorm.DB`，数据访问层。

### 预留的空目录

`internal/model/`、`internal/router/`、`internal/middware/`（注意：拼写错误，应为 `middleware`）、`pkg/`、`deployments/`、`docs/`、`scripts/`、`test/`、`utils/`

## 配置

启动时读取 `config/config.yaml`。Dockerfile 会将其复制到镜像中——不同环境建议挂载替换配置文件或改用环境变量覆盖。

## Docker

```bash
docker build -t bi-api .
docker run -p 8080:8080 bi-api
```

没有 `docker-compose.yml`。配置文件中的 MySQL 地址为 `127.0.0.1:3306`，在容器内部无法解析到宿主机——本地 Docker 开发时需改用 `host.docker.internal` 或编写带 MySQL 服务的 compose 文件。
