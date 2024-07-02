package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

func SHA(c *gin.Context) {
	password := c.Query("password")
	if password == "" {
		password = c.PostForm("password")
	}
	// 创建了一个新的 SHA-1 哈希生成器实例
	Sha1 := sha1.New()
	// 将提供的密码转换为字节切片
	Sha1.Write([]byte(password))
	// Sha1.Sum(nil): 这个方法完成哈希计算并将结果作为字节切片返回
	// 将哈希的字节切片表示转换为十六进制字符串表示
	c.Set("password", hex.EncodeToString(Sha1.Sum(nil)))
	c.Next()
}
