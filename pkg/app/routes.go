package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV2 = "/api/v2"

	// specification routes
	a.Router.HandleFunc(apiV2+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")

	// user routes
	a.Router.HandleFunc(apiV2+"/", a.getAllTickets).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
