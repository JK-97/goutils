# Go Utils

Golang 工具库

## 使用 SDK

### 安装 SDK

创建对应的 Golang 项目，并初始化为 git 仓库或克隆已有的项目。

使用 `go mod`

```bash
go mod init
```

添加 git 子模块

```bash
git submodule add http://gitlab.jiangxingai.com/applications/base-modules/internal-sdk/go-utils.git
```

替换依赖为本地版本

```bash
go mod edit -require gitlab.jiangxingai.com/applications/base-modules/internal-sdk/go-utils@v0.0.0-00010101000000-000000000000
go mod edit -replace gitlab.jiangxingai.com/applications/base-modules/internal-sdk/go-utils=./go-utils
```

### 更新使用的 SDK 版本

```bash
git submodule update --remote
```

## Logger

日志库
