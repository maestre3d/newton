package infrastructure

import "go.uber.org/zap"

// NewZapLogger creates a new Uber's Zap logger depending on the development stage
func NewZapLogger(cfg Configuration) (*zap.Logger, error) {
	if cfg.IsProd() {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}
