# config-file

`config-file` 是 `config` 模块的 `file` 驱动。

## 安装

```bash
go get github.com/infrago/config@latest
go get github.com/infrago/config-file@latest
```

## 接入

```go
import (
    _ "github.com/infrago/config"
    _ "github.com/infrago/config-file"
    "github.com/infrago/infra"
)

func main() {
    infra.Run()
}
```

## 配置示例

```toml
[config]
driver = "file"
```

## 公开 API（摘自源码）

- `func (d *FileConfigDriver) Load(params Map) (Map, error)`

## 排错

- driver 未生效：确认模块段 `driver` 值与驱动名一致
- 连接失败：检查 endpoint/host/port/鉴权配置
