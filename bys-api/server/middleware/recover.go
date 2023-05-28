package middleware

import (
	"bys/pkg/ulog"
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				msg := fmt.Sprintf("server panic: err Type = %+v, err = %+v, stack=%v", reflect.TypeOf(r), r,
					string(debug.Stack()))
				ulog.Error(msg)

				c.JSON(http.StatusInternalServerError, nil)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
