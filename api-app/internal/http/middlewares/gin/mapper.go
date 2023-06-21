package gin

import (
	rest "github.com/YReshetko/it-learning-platform/api-app/internal/http"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Wrap[Request any, Response any](fn rest.HandlerFunc[Request, Response], redirect string, logger *logrus.Entry) func(*gin.Context) {
	return func(ginCtx *gin.Context) {
		var rq Request
		_, ok := any(rq).(models.Empty)

		if !ok {
			if err := ginCtx.ShouldBindJSON(&rq); err != nil {
				logger.WithError(err).Error("parse body error")
				ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		ctx := rest.Context{
			GinCtx: ginCtx,
		}

		rs, status := fn(ctx, rq)
		if status.StatusCode == http.StatusUnauthorized {
			logger.WithField("redirect_to", redirect).Info("redirecting as unauthorized")
			ginCtx.Redirect(http.StatusFound, redirect)
			return
		}

		if status.Error == nil {
			_, ok = any(rs).(models.Empty)
			if !ok {
				ginCtx.JSON(status.StatusCode, rs)
			}
			return
		}
		logger.WithField("error_status", status.StatusCode).WithError(status.Error).Error(status.Message)
		ginCtx.JSON(status.StatusCode, gin.H{"error": status.Message})
	}
}
