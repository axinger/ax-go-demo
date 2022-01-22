2.3 MAC主机编译Widows,linux客户端
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o abc-demo-linux main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o abc-demo-mac main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o abc-demo-windows.exe main.go
```