package infrastructure

import (
	"github.com/spf13/viper"
)

// Configuration kernel/global configuration using OS environment variables if prod and yaml config file for the rest
// stages
type Configuration struct {
	Stage       string
	Version     string
	DynamoTable string
}

func init() {
	viper.SetDefault("newton.stage", DevStage)
	viper.SetDefault("newton.version", "1.0.0")
	viper.SetDefault("newton.dynamo.table", "newton-books-dev")
}

// NewConfiguration creates a Configuration with default configs or from sources
func NewConfiguration() Configuration {
	return Configuration{
		Stage:       viper.GetString("newton.stage"),
		Version:     viper.GetString("newton.version"),
		DynamoTable: viper.GetString("newton.dynamo.table"),
	}
}

const (
	// ProdStage Production deployment stage
	ProdStage = "prod"
	// Dev Development deployment stage
	DevStage = "dev"
)

// IsProd returns if current config stage is in production stage
func (c Configuration) IsProd() bool {
	return c.Stage == ProdStage
}
