package serverlessplus

import (
	"net/http"
	"os"

	"github.com/ipfans/serverlessplus/aliyun"
	"github.com/ipfans/serverlessplus/tencent"
)

// Start 启动Serverless服务。
// 不支持平台会绑定至SERVERLESS_ADDR环境变量指定的地址或本地9999端口。
func Start(handler http.Handler) (err error) {
	switch DetectPlatform() {
	case PlatformAliyun:
		err = aliyun.Start(handler)
	case PlatformTencent:
		err = tencent.Start(handler)
	case PlatformUnsupport:
		addr := os.Getenv("SERVERLESS_ADDR")
		if addr == "" {
			addr = "127.0.0.1:9999"
		}
		err = http.ListenAndServe(addr, handler)
	}
	return
}
