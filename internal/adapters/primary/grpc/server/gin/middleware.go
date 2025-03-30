package gin

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recoverer is a middleware that recovers from panics and returns a JSON error response.
func Recoverer() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rvr := recover(); rvr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("%v", rvr),
					"stack": string(debug.Stack()),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
