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
			fmt.Println("parse body error: ", err)
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("body: %+v\n", rq)
		ctx := rest.Context{
			GinCtx: ginCtx,
		}

		rs, status := fn(ctx, rq)
		if status.StatusCode == http.StatusUnauthorized {
			url := "http://localhost:8081/realms/it-academy/protocol/openid-connect/auth?response_type=token&scope=openid&client_id=academy&redirect_uri=http://localhost:8080"
			//url := "http://localhost:8080/auth?response_type=token&scope=openid&client_id=academy&redirect_uri=http://localhost:8080"
			fmt.Println("REDIRECT TO: ", url)
			/*redirectRequest, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				fmt.Println("unable to create redirect request", err)
			}
			redirectRequest.Header.Add("Access-Control-Allow-Origin", "*")
			ginCtx.Render(http.StatusUnauthorized, render.Redirect{
				Code:     http.StatusFound,
				Location: url,
				Request:  redirectRequest,
			})*/
			ginCtx.Redirect(http.StatusFound, url)
			return
		}

		if status.Error == nil {
			ginCtx.JSON(status.StatusCode, rs)
			return
		}
		fmt.Println(status.Error)
		ginCtx.JSON(status.StatusCode, gin.H{"error": status.Message})
	}
}
