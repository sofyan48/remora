package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/remora/src/app/v1/api/playbook/entity"
	"github.com/sofyan48/remora/src/app/v1/api/playbook/service"
	"github.com/sofyan48/remora/src/app/v1/utility/rest"
)

// PlaybookController types
type PlaybookController struct {
	Service service.PlaybookServiceInterface
}

// PlaybookControllerHandler ...
func PlaybookControllerHandler() *PlaybookController {
	return &PlaybookController{
		Service: service.PlaybookServiceHandler(),
	}
}

// PlaybookControllerInterface ...
type PlaybookControllerInterface interface {
	Playbook(context *gin.Context)
}

// Playbook params
// @contex: gin Context
func (ctrl *PlaybookController) Playbook(context *gin.Context) {
	params := &entity.PlaybookParameters{}
	if err := context.ShouldBind(&params); err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	err := ctrl.Service.PlaybookService(params.Apps, params.Connection, params.Inventory)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, "Playbook Executor")
	return
}
