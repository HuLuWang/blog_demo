package middleware

import (
	"github.com/HuLuWang/blog_demo/pkg/app"
	"github.com/HuLuWang/blog_demo/pkg/errcode"
	"github.com/HuLuWang/blog_demo/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		
		c.Next()
	}
}

