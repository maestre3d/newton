package httputil

import (
	"strconv"

	"github.com/maestre3d/newton/internal/infrastructure"
)

// NewVersioning generates a valid URI with given version
func NewVersioning(cfg infrastructure.Configuration) string {
	if major := cfg.MajorVersion(); major > 0 {
		return "/v" + strconv.Itoa(major)
	}
	return "/" + cfg.ReleaseStage()
}
