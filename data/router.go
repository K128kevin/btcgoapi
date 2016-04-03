package data

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

var root = "/api/data"

var routes = utils.Routes {
	utils.Route {
		"Index",
		"GET",
		root + "/prices",
		GetPrices,
	},
	utils.Route {
		"Index",
		"GET",
		root + "/prediction",
		GetPredictions,
	},
	utils.Route {
		"Index",
		"POST",
		root + "/prices",
		AddPrice,
	},
	utils.Route {
		"Index",
		"POST",
		root + "/prediction",
		AddPrediction,
	},
}