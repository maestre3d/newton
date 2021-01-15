package infrastructure

import (
	"github.com/spf13/viper"
)

// Configuration kernel/global configuration using OS environment variables if prod and yaml config file for the rest
// stages
type Configuration struct {
	Stage       string
	Version     string
	AdminEmail  string
	DynamoTable string
}

func init() {
	viper.SetDefault("newton.stage", DevStage)
	viper.SetDefault("newton.version", "1.0.0")
	viper.SetDefault("newton.admin_email", "luis.alonso.16@hotmail.com")
	viper.SetDefault("newton.dynamo.table", "newton-books-dev")
}

const (
	// ProdStage Production deployment stage
	ProdStage = "prod"
	// DevStage Development deployment stage
	DevStage = "dev"
)

// NewConfiguration creates a Configuration with default configs or from sources
func NewConfiguration() Configuration {
	return Configuration{
		Stage:       viper.GetString("newton.stage"),
		Version:     viper.GetString("newton.version"),
		AdminEmail:  viper.GetString("newton.admin_email"),
		DynamoTable: viper.GetString("newton.dynamo.table"),
	}
}

// IsProd returns if current config stage is in production stage
func (c Configuration) IsProd() bool {
	return c.Stage == ProdStage
}
