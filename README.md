# Delete failed action

English | [中文](./README_zh.md)

Delete your fetal action.

# Usage

```shell
export GH_TOKEN="your token"
./delete-failed-action -u "xchacha20-poly1305" -r "delete-failed-action" -t $GH_TOKEN -w "build.yml"
```

# Develop

## Build

```shell
go build -v . -trimpath -ldflags="-w -s -buildid="
```