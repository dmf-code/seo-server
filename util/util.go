package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"syscall"
	"time"
)

func Md5(str []byte) string {
	hash := md5.New()

	hash.Write(str)

	md5Code := fmt.Sprintf("%x", hex.EncodeToString(hash.Sum(nil)))
	return md5Code
}

func PathExists(path string) (isExist bool) {
	if _, err := os.Lstat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Expire(path string, sec int64) bool {
	fileInfo, _ := os.Stat(path)

	var lastWriteTime int64


	// windows
	fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	nanoseconds := fileSys.LastWriteTime.Nanoseconds() // 返回的是纳秒
	lastWriteTime = int64(nanoseconds/1e9) //秒

	// linux
	//fileSys := fileInfo.Sys().(*syscall.Stat_t)
	//lastWriteTime = int64(fileSys.Mtim.Sec)


	now := time.Now().Unix()

	if now - lastWriteTime >= sec {
		return true
	}

	return false
}
