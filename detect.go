package serverlessplus

import (
	"os"
)

const (
	// PlatformUnsupport 平台暂不支持
	PlatformUnsupport = iota
	// PlatformAliyun 阿里云FC
	PlatformAliyun
	// PlatformTencent 腾讯云SCF
	PlatformTencent
)

// DetectPlatform 用于检测当前serverless平台，目前支持检测阿里云与腾讯云。
func DetectPlatform() int {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" {
		return PlatformTencent
	} else if os.Getenv("FC_SERVER_PORT") != "" {
		return PlatformAliyun
	}
	return PlatformUnsupport
}
