package controller

import (
	"fmt"

	"MANAGEMENT-2.0/store"
)

type Server struct {
	PostgressDb store.StoreOperation
}

func (g *Server) NewServer(store store.Postgres) {
	g.PostgressDb = &store
	g.PostgressDb.NewStore()

	fmt.Printf("Server : %+v\n", g.PostgressDb)
}


type ServerOperation interface{
	NewServer(store store.Postgres)
}