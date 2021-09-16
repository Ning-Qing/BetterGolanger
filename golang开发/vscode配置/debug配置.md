```json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "example debug",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "service/example/main.go", //启动路劲
            "args": ["--registry=etcd","--registry_address=127.0.0.1:2379"], // 启动参数
            "dlvFlags": ["--check-go-version=false"], // 调试工具dlv参数，禁用go版本检查
            "env": {
                "jaeger_address":"127.0.0.1:6831" // 环境变量
            },
            "showLog":true // 显示日志
        }
    ]
}
```