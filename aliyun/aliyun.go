package aliyun

import (
	"net/http"
	"os"
)

func wrapHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-Request-Id", r.Header.Get("X-FC-Request-Id"))
	})
}

// Start 启动阿里云函数计算服务
func Start(handler http.Handler) (err error) {
	return http.ListenAndServe(":"+os.Getenv("FC_SERVER_PORT"), wrapHandler(handler))
}
