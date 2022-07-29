package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GaurKS/book-api/rate-limiter-pkg/rateLimiter"
	"github.com/gin-gonic/gin"
)

var _ipMap = make(map[string]*rateLimiter.Bucket) // used for checking IP spamming
const MAX_LIMIT = 2                               // max limit to api call

func RateLimit(c *gin.Context) {
	ip := c.ClientIP()
	fmt.Println("Current client IP: ", ip)
	requestBucket := GetBucket(ip)
	flag, tryAfter, rem := requestBucket.IsAllowed(1)
	if !flag {
		c.Header("content-type", "application/json")
		c.Header("Retry-After", tryAfter.String())
		c.AbortWithStatusJSON(
			http.StatusTooManyRequests,
			gin.H{
				"Message": "Request limit exhausted",
			},
		)
	} else {
		c.Header("X-RateLimit-Limit", strconv.Itoa(MAX_LIMIT))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(rem, 10))
	}
	c.Next()
}

func GetBucket(ip string) *rateLimiter.Bucket {
	if _ipMap[ip] == nil {
		_ipMap[ip] = rateLimiter.NewBucket(MAX_LIMIT, ip)
	}
	return _ipMap[ip]
}
