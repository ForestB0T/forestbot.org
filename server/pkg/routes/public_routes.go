package routes

import (
	"net/http"

	"github.com/febzey/forestbot-api/pkg/controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func PublicRoutes(router *mux.Router) {
	r := controllers.Routes{}

	var routes = []Route{
		{
			Method:      http.MethodGet,
			Pattern:     "/test/{param}",
			HandlerFunc: r.GetTest,
		},
	}

	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
}
