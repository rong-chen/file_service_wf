package utils

import (
	"runtime"
)

// GetOsPath 获取当前os系统
func GetOsPath(path string) string {
	if runtime.GOOS == "windows" {
		return "." + path
	}
	return path
}
