package api

import (
	"MANAGEMENT-2.0/model"
	"MANAGEMENT-2.0/util"
	"github.com/gin-gonic/gin"
)

func (api ApiRouts) UserRoutes(routes *gin.Engine) {
	group := routes.Group("/user")
	{
		group.POST("/create", api.CreatUser)
	}
}

func (api ApiRouts) CreatUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreatUser, " creating new user", nil)

	api.Server.CreatUser(ctx)
}
