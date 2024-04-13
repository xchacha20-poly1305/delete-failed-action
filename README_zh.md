# 删除失败的 actions！

[English](./README.md) | 中文

调试 CI 到爆炸？一键清除您失败的痕迹！

# 使用

```shell
export GH_TOKEN="your token"
./delete-failed-action -u "xchacha20-poly1305" -r "delete-failed-action" -t $GH_TOKEN -w "build.yml"
```

# 开发

## 构建

```shell
go build . -v -trimpath -ldflags="-w -s -buildid="
```
