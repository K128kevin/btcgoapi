package tradesim

import (
	"github.com/gorilla/mux"
	"btc_git_1.6/utils"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)

	}
	return router
}

var root = "/api/tradesim"

var routes = utils.Routes {
	utils.Route {
		"GetTrade",
		"GET",
		root + "/{tradeId}",
		GetTrade,
	},
}