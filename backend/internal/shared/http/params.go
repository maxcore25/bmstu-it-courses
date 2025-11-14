package httphelper

import "github.com/gin-gonic/gin"

func ShouldInclude(c *gin.Context, key string) bool {
	return c.Query("include") == key
}
