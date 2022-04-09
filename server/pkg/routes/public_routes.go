package routes

import (
	"database/sql"
	"net/http"

	"github.com/febzey/forestbot-api/pkg/controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func PublicRoutes(router *mux.Router, db *sql.DB) {
	r := controllers.Routes{
		DB: db,
	}

	var routes = []Route{
		{
			Method:      http.MethodGet,
			Pattern:     "/test/{user}/{db}",
			HandlerFunc: r.Test,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/playtime/{id}/",
			HandlerFunc: r.GetPlaytime,
		},
	}

	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
}
