package logger

import "go.uber.org/zap"

func New(env string) *zap.Logger {
	if env == "dev" {
		l, _ := zap.NewDevelopment()
		return l
	}
	l, _ := zap.NewProduction()
	return l
}
