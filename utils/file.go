package utils

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/md4"
	"io"
	"os"
	"path/filepath"
)

const chunkSize = 9728000 // 9.28MB
// 计算 ED2K 哈希
func ed2kHash(filePath string) (string, int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	var chunkHashes []byte
	buf := make([]byte, chunkSize)
	hasher := md4.New()

	// 逐块计算 MD4 哈希
	for {
		n, err := file.Read(buf)
		if n > 0 {
			chunkHash := md4.New()
			chunkHash.Write(buf[:n])
			chunkHashes = append(chunkHashes, chunkHash.Sum(nil)...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", 0, err
		}
	}

	// 计算最终 ED2K 哈希
	if len(chunkHashes) == md4.Size {
		return hex.EncodeToString(chunkHashes), fileSize(filePath), nil
	}
	hasher.Write(chunkHashes)
	return hex.EncodeToString(hasher.Sum(nil)), fileSize(filePath), nil
}

// 获取文件大小
func fileSize(filePath string) int64 {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return info.Size()
}

// GenerateED2K 生成 ED2K 链接
func GenerateED2K(filePath string) (string, error) {
	hash, size, err := ed2kHash(filePath)
	if err != nil {
		return "", err
	}
	fileName := filepath.Base(filePath)
	return fmt.Sprintf("ed2k://|file|%s|%d|%s|/", fileName, size, hash), nil
}
