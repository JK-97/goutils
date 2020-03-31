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
git submodule add github.com/JK-97/go-utils.git
```

替换依赖为本地版本

```bash
go mod edit -require github.com/JK-97/go-utils.git
go mod edit -replace github.com/JK-97/go-utils.git=./go-utils
```

### 更新使用的 SDK 版本

```bash
git submodule update --remote
```

## Logger

日志库
