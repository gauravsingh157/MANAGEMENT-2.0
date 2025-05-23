package main

import (
	"MANAGEMENT-2.0/api"
	"MANAGEMENT-2.0/controller"
)

func main() {
	api := api.ApiRouts{}
	api.Startapp(controller.Server{})

	//fmt.Printf("Main server: %+v\n", api)
}
