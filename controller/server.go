package controller

import (
	"fmt"

	"MANAGEMENT-2.0/model"
	"MANAGEMENT-2.0/store"
	"MANAGEMENT-2.0/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	PostgressDb store.StoreOperation
}

func (s *Server) NewServer(pgstore store.Postgres) {

	util.SetLogger()
	util.Logger.Infof("Logger setup Done ...")

	s.PostgressDb = &pgstore

	err := s.PostgressDb.NewStore()
	if err != nil {
		util.Logger.Errorf("Error while creating store connection ,err =%v\n", err)

	} else {
		util.Logger.Infof("Connecting with store\n ")
		util.Log(model.LogLevelInfo, model.ControllerPackage, "NewServer", "Connecting with store", nil)

	}
	fmt.Printf("Server : %+v\n", s.PostgressDb)
}

type ServerOperation interface {
	NewServer(store store.Postgres)

	CreatUser(ctx *gin.Context)
}
