```shell
# 安装特定的go版本包装器
go get golang.org/dl/go<version>
go get golang.org/dl/go1.13
# 安装
go<version> download
go1.13 download

# go1.13 env 配置
go1.13 env -w GO111MODULE="on"
go1.13 env -w GOPROXY="https://goproxy.io,direct"
```