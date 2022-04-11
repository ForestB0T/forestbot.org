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
			Pattern:     "/playtime/{user}/{server}",
			HandlerFunc: r.GetPlaytime,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/kd/{user}/{server}",
			HandlerFunc: r.GetKD,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/joins/{user}/{server}",
			HandlerFunc: r.GetJoins,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/lastseen/{user}/{server}",
			HandlerFunc: r.GetLastSeen,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/joindate/{user}/{server}",
			HandlerFunc: r.GetJoinDate,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/user/{user}/{server}",
			HandlerFunc: r.GetAllUserStats,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/messagecount/{user}/{server}",
			HandlerFunc: r.GetMessageCount,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/quote/{user}/{server}",
			HandlerFunc: r.GetQuote,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/lastdeath/{user}/{server}",
			HandlerFunc: r.GetLastDeath,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/tab/{server}",
			HandlerFunc: r.GetTablist,
		},
	}

	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
}
