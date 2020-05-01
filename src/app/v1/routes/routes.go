package routes

import (
	"github.com/gin-gonic/gin"
	health "github.com/sofyan48/remora/src/app/v1/api/health/controller"
	playbook "github.com/sofyan48/remora/src/app/v1/api/playbook/controller"

	"github.com/sofyan48/remora/src/middleware"
)

// VERSION ...
const VERSION = "v1"

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Health     health.HealthControllerInterface
	Playbook   playbook.PlaybookControllerInterface
}

// V1RouterLoaderHandler ...
func V1RouterLoaderHandler() *V1RouterLoader {
	return &V1RouterLoader{
		Health:   health.HealthControllerHandler(),
		Playbook: playbook.PlaybookControllerHandler(),
	}
}

// V1RouterLoaderInterface ...
type V1RouterLoaderInterface interface {
	V1Routes(router *gin.Engine)
}

// V1Routes Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Routes(router *gin.Engine) {
	rLoader.initDocs(router)
	rLoader.initHealth(router)
	rLoader.initPlaybook(router)
}
