package main

import (
	"MANAGEMENT-2.0/api"
	"MANAGEMENT-2.0/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	api := api.ApiRouts{}

	controller := controller.Server{}
	routes := gin.Default()
	api.Startapp(routes, controller)
	routes.Run(":8000")
	//fmt.Printf("Main server: %+v\n", api)
}
