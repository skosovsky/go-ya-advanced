package main

import "go.uber.org/zap"

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer logger.Sync() //nolint:errcheck // example

	const url = "https://www.google.com"

	sugar := logger.Sugar()

	sugar.Infow("Failed to fetch URL",
		"url", url)

	sugar.Infof("Failed to fetch URL: %s", url)
	sugar.Errorf("Failed to fetch URL: %s", url)

	plain := sugar.Desugar()

	plain.Info("Hello, Go!")
	plain.Warn("Simple warning message")
	plain.Error("Failed to fetch URL", zap.String("url", url))
}
