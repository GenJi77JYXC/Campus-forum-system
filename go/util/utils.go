package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

var (
	avatarURLTemplate = "https://robohash.org/%s?set=set%d"
	// avatarGroupNum    = []int{1, 2, 3, 4}
)

// 返回随机头像url
// thanks to https://robohash.org
func RandomAvatarURL(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)                                                     // {Sum(b []byte) []byte}  Sum 将当前哈希附加到 b 并返回生成的切片。（它不会更改基础哈希状态。）
	return fmt.Sprintf(avatarURLTemplate, hex.EncodeToString(cipherStr), str[0]%4+1) // EncodeToString 返回 src 的十六进制编码。
}

// 截取字符串
func SubString(str string, start int, end int) string {
	if start >= len(str) || end >= len(str) || start > end || end < 0 {
		return ""
	}
	return string([]rune(str)[start:end])
}

// 获取当前时间戳
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6 //时间戳（毫秒）
}

// Timestamp返回t的时间戳
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6 //时间戳（毫秒）
}

// 比较两个int64值大小，返回较小值
func MinInt64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// 比较两个int64值大小，返回较大值
func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// 比较两个int值大小，返回较小值
func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// 比较两个int值大小，返回较大值
func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
