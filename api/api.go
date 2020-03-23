package api

import (
	"github.com/gorilla/mux"
	"github.com/nagarjun226/configmgr/controller"
)

// API - API object. A nice abstraction around the router. Maybe can be used in the future, not of much use now
// useful if you have to store login credentials?
type API struct {
}

// Router - returns a router to direct API calls for a controller
func (api *API) Router(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getconfig/{service}", c.ServeConfig).Methods("GET")
	return r
}
