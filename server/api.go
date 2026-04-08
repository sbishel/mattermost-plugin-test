package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost/server/public/plugin"
)

// initRouter initializes the HTTP router for the plugin.
func (p *Plugin) initRouter() *mux.Router {
	router := mux.NewRouter()

	// Dialog handlers (no auth middleware — called by Mattermost server)
	dialogRouter := router.PathPrefix("/dialog").Subrouter()
	dialogRouter.HandleFunc("/submit-with-validation", p.handleDialogSubmitWithValidation).Methods(http.MethodPost)
	dialogRouter.HandleFunc("/submit-confirm", p.handleDialogSubmitConfirm).Methods(http.MethodPost)
	dialogRouter.HandleFunc("/submit-generic", p.handleDialogSubmitGeneric).Methods(http.MethodPost)
	dialogRouter.HandleFunc("/error", p.handleDialogWithError).Methods(http.MethodPost)
	dialogRouter.HandleFunc("/field-refresh", p.handleDialogFieldRefresh).Methods(http.MethodPost)
	dialogRouter.HandleFunc("/multistep", p.handleDialogMultistep).Methods(http.MethodPost)
	dialogRouter.HandleFunc("/roles", p.handleDynamicRoles).Methods(http.MethodPost)

	return router
}

// ServeHTTP handles HTTP requests to the plugin.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}
