package api

import (
	"MANAGEMENT-2.0/controller"
	"MANAGEMENT-2.0/store"
)

type ApiRouts struct {
	Server controller.ServerOperation
}

func (api *ApiRouts) Startapp(server controller.Server) {
	api.Server = &server

	api.Server.NewServer(store.Postgres{})
	

}
