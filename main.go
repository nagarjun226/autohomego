package main

import (
	"net/http"

	"github.com/nagarjun226/configmgr/api"
	"github.com/nagarjun226/configmgr/controller"

	"github.com/nagarjun226/configmgr/domain"
)

func main() {
	c := domain.Config{}

	autoLoader := domain.ConfigAutoLoader{
		Config:   &c,
		Location: "config.json",
		Rr:       3,
	}
	go autoLoader.Run()

	controller := controller.Controller{
		Config: &c,
	}

	api := api.API{}

	r := api.Router(&controller)
	http.ListenAndServe(":8080", r)
}
