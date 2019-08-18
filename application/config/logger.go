package config

import "go.uber.org/zap"

func NewLogger(logPath string) (*zap.Logger, error) {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.OutputPaths = []string{logPath}
	logConfig.DisableCaller = true
	return logConfig.Build()
}
