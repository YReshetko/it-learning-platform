package gin

import (
	"fmt"
	rest "github.com/YReshetko/it-academy-cources/api-app/internal/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Wrap[Request any, Response any](fn rest.HandlerFunc[Request, Response]) func(*gin.Context) {
	return func(ginCtx *gin.Context) {
		var rq Request
		if err := ginCtx.ShouldBindJSON(&rq); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := rest.Context{
			GinCtx: ginCtx,
		}

		rs, status := fn(ctx, rq)
		if status.Error != nil {
			fmt.Println(status.Error)
			ginCtx.JSON(status.StatusCode, gin.H{"error": status.Message})
			return
		}
		ginCtx.JSON(status.StatusCode, rs)
	}
}
