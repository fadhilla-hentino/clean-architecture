package controller

import (
	"net/http"

	"fadhilla-hentino/clean-architecture/model"
	"fadhilla-hentino/clean-architecture/service"
	"fadhilla-hentino/clean-architecture/utils"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Controller struct {
	svc    service.FriendService
	logger *zap.Logger
}

func NewController(svc service.FriendService, logger *zap.Logger) *Controller {
	return &Controller{
		svc:    svc,
		logger: logger,
	}
}

func (c *Controller) Route(r *httprouter.Router) {
	r.POST("/friends", c.RequestController)
	r.GET("/friends/users/:userID", c.RequestController)
}

func (c *Controller) RequestController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var req *model.FriendRequestReq
	if err := utils.ReadEntity(r, &req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "unable to parse request body")
		return
	}

	if err := c.svc.Request(r.Context(), req.UserID, req.FriendID); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "unable to request friend")
		return
	}

	utils.WriteOKNoContentResponse(w)
}
