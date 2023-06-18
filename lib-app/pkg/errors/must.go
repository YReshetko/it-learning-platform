package errors

import (
	"github.com/sirupsen/logrus"
	"os"
)

func MustExitAppErrorHandler[T any](logger *logrus.Entry) func(val T, err error) T {
	return func(val T, err error) T {
		if err != nil {
			logger.WithError(err).Error("crashing the app")
			os.Exit(1)
		}
		return val
	}
}
