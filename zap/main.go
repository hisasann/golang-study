package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	url := "http://example.com"

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	suger := logger.Sugar()
	suger.Infow("Hello")
	suger.Infof("Hello, %s", "世界")
}
