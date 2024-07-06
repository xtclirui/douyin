package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

// 符合最小权限原则：服务端加密密码后，即使服务端被入侵，攻击者也无法直接获取用户的明文密码。
// 密码存储安全：服务端需要存储用户密码的哈希值而不是明文密码
// 确保即使数据库被盗，攻击者也无法轻易破解出原始密码

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
