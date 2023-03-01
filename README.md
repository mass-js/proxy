# proxy  在命令行中输入 go run proxy.go -listen=:8080 -target=http://example.com:80 启动代理服务器。其中 -listen 参数指定代理服务器监听的地址和端口， -target 参数指定代理服务器需要将请求转发到的目标服务器的地址和端口。
在浏览器中访问 http://localhost:8080 即可通过代理服务器访问目标服务器。如果目标服务器返回的内容中包含
