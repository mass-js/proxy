package main

import (
    "flag"
    "log"
    "net"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    // 定义命令行参数
    listenAddr := flag.String("listen", ":8080", "代理服务器监听地址")
    targetURL := flag.String("target", "", "代理目标URL")
    flag.Parse()

    // 解析目标URL
    target, err := url.Parse(*targetURL)
    if err != nil {
        log.Fatal("无效的目标URL:", err)
    }

    // 创建反向代理器
    proxy := httputil.NewSingleHostReverseProxy(target)

    // 创建HTTP处理函数
    handler := func(w http.ResponseWriter, r *http.Request) {
        // 设置X-Forwarded-For头部，以便目标服务器获取真实客户端IP地址
        r.Header.Set("X-Forwarded-For", r.RemoteAddr)

        // 转发请求到目标服务器
        proxy.ServeHTTP(w, r)
    }

    // 创建TCP监听器
    l, err := net.Listen("tcp", *listenAddr)
    if err != nil {
        log.Fatal("无法监听:", err)
    }

    // 创建HTTP服务器并开始监听
    log.Printf("代理服务器已启动，监听地址：%v，目标地址：%v", *listenAddr, *targetURL)
    err = http.Serve(l, http.HandlerFunc(handler))
    if err != nil {
        log.Fatal("无法启动HTTP服务器:", err)
    }
}
