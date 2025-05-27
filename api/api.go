package api

import (
	"MANAGEMENT-2.0/controller"
	"MANAGEMENT-2.0/store"
	"github.com/gin-gonic/gin"
)

type ApiRouts struct {
	Server controller.ServerOperation
}

func (api *ApiRouts) Startapp(routes *gin.Engine, server controller.Server) {

	api.Server = &server

	api.Server.NewServer(store.Postgres{})

	api.UserRoutes(routes)

}
