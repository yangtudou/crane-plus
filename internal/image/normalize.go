package image

import "strings"

// NormalizeRegistry 返回规范化后的 registry。
//
// Docker Hub:
// nginx
// library/nginx
// docker.io/library/nginx
//
// 最终统一为:
//
// docker.io
func NormalizeRegistry(registry string) string {
	if registry == "" {
		return "docker.io"
	}

	return registry
}

// IsDockerHub 判断是否 Docker Hub。
func IsDockerHub(registry string) bool {
	return strings.TrimSpace(registry) == "docker.io"
}
