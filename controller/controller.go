package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nagarjun226/configmgr/domain"
)

// Controller - Handles Serving up config information on request
type Controller struct {
	Config *domain.Config // holds the address of the config object
}

// ServeConfig - Handler function for serving configs
func (c *Controller) ServeConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") // Set the header to tell the requester what format the daya is in

	vars := mux.Vars(r) // Get the parameteres in the request. In this case, the service name only
	service, ok := vars["service"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error")
	}

	config, err := c.Config.GetConfig(service)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err))
	}

	rsp, err := json.Marshal(&config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err))
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(rsp))
}
