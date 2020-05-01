package routes

import (
	"github.com/gin-gonic/gin"
)

// PLAYBOOKROUTES ...
const PLAYBOOKROUTES = VERSION + "/playbook"

func (rLoader *V1RouterLoader) initPlaybook(router *gin.Engine) {
	group := router.Group(PLAYBOOKROUTES)
	group.GET("", rLoader.Playbook.Playbook)
}
