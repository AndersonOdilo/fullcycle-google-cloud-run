package main

import (
	"fmt"

	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/configs"
	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/api"
	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/infra/web"
	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/infra/web/webserver"
)

func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	locationRepository := api.NewLocationRepository();
	tempRepository := api.NewTempRepository(*configs);

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webTempHandler := web.NewWebTempHandler(locationRepository, tempRepository);
	webserver.AddHandler("GET /temp/{cep}", webTempHandler.Get)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start()

}