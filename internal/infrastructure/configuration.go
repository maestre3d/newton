package infrastructure

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// Configuration kernel/global configuration using OS environment variables if prod and yaml config file for the rest
// stages
type Configuration struct {
	Application string
	Stage       string
	Version     string
	HTTPAddress string
	HTTPPort    int
	AdminEmail  string
	DynamoTable string
	AWSEventBus string
}

func init() {
	viper.SetDefault("newton.application", "org.neutrinocorp.newton")
	viper.SetDefault("newton.stage", DevStage)
	viper.SetDefault("newton.version", "1.0.0")
	viper.SetDefault("newton.http", "")
	viper.SetDefault("newton.http.port", 8081)
	viper.SetDefault("newton.admin_email", "luis.alonso.16@hotmail.com")
	viper.SetDefault("newton.dynamo.table", "newton-authors-dev")
	viper.SetDefault("newton.eventbridge.bus", "newton-dev")
}

const (
	// ProdStage Production deployment stage
	ProdStage = "prod"
	// DevStage Development deployment stage
	DevStage = "dev"
)

// NewConfiguration creates a Configuration with default configs or from sources
func NewConfiguration() Configuration {
	viper.AutomaticEnv()
	return Configuration{
		Application: viper.GetString("newton.application"),
		Stage:       viper.GetString("newton.stage"),
		Version:     viper.GetString("newton.version"),
		HTTPAddress: viper.GetString("newton.http"),
		HTTPPort:    viper.GetInt("newton.http.port"),
		AdminEmail:  viper.GetString("newton.admin_email"),
		DynamoTable: viper.GetString("newton.dynamo.table"),
		AWSEventBus: viper.GetString("newton.eventbridge.bus"),
	}
}

// IsProd returns if current config stage is in production stage
func (c Configuration) IsProd() bool {
	return c.Stage == ProdStage
}

// MajorVersion returns the current major version
func (c Configuration) MajorVersion() int {
	major, err := strconv.Atoi(strings.Split(c.Version, ".")[0])
	if err != nil {
		return 0
	}

	return major
}

// ReleaseStage returns the current release stage
func (c Configuration) ReleaseStage() string {
	stage := strings.Split(c.Version, "-")
	if len(stage) < 2 {
		return ""
	}

	return stage[1]
}
