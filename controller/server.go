package controller

import (
	"fmt"

	"MANAGEMENT-2.0/model"
	"MANAGEMENT-2.0/store"
	"MANAGEMENT-2.0/util"
)

type Server struct {
	PostgressDb store.StoreOperation
}

func (g *Server) NewServer(store store.Postgres) {
	util.SetLogger()
	util.Logger.Infof("Logger setup Done ...")
	g.PostgressDb = &store
	err := g.PostgressDb.NewStore()
	if err != nil {
		util.Logger.Errorf("Error while creating store connection ,err =%v\n", err)

	} else {
		util.Logger.Infof("Connecting with store\n ")
		util.Log(model.LogLevelInfo, model.Controller, "NewServer", "Connecting with store", nil)

	}
	fmt.Printf("Server : %+v\n", g.PostgressDb)
}

type ServerOperation interface {
	NewServer(store store.Postgres)
}
