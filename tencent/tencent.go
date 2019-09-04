package tencent

import (
	"net"
	"net/http"

	"github.com/akutz/memconn"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

const (
	serverMemAddr = "ServerlessPlus:80"
)

// Start 启动腾讯云函数计算服务器
func Start(handler http.Handler) (err error) {
	// 启动内存WebServer监听服务
	var lis net.Listener
	lis, err = memconn.Listen("memu", serverMemAddr)
	if err != nil {
		return
	}
	go http.Serve(lis, handler)

	cloudfunction.Start(toHTTPRequest)
	return
}
