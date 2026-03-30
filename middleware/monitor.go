package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PrometheusMiddleware Prometheus监控中间件
// 用于统计API的请求量、响应时间、状态码等指标
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 处理请求
		c.Next()

		// 请求完成后记录指标
		duration := time.Since(start).Milliseconds()
		statusCode := c.Writer.Status()

		// 这里可以记录到日志或发送到监控系统
		// 示例：打印到控制台
		if statusCode >= 400 {
			// 错误请求
			println("[ERROR]", method, path, "status:", statusCode, "duration:", duration+"ms")
		} else {
			// 正常请求
			println("[INFO]", method, path, "status:", statusCode, "duration:", duration+"ms")
		}
	}
}

// RecoveryMiddleware 自定义错误恢复中间件
// 用于捕获panic并记录错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				println("[PANIC]", "error:", err, "path:", c.Request.URL.Path)
				c.JSON(500, gin.H{
					"code":    500,
					"message": "服务器内部错误",
					"error":   err,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// HealthCheckMiddleware 健康检查中间件
// 用于统计服务健康状态
var HealthStatus = struct {
	TotalRequests  int64
	ErrorRequests  int64
	LastCheckTime  time.Time
	DatabaseStatus string
	RedisStatus    string
}{
	TotalRequests:  0,
	ErrorRequests:  0,
	LastCheckTime:  time.Now(),
	DatabaseStatus: "unknown",
	RedisStatus:    "unknown",
}
