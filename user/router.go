package user

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

var root = "/api/users"

var routes = utils.Routes {
	utils.Route {
		"GetUserInfo",
		"GET",
		root + "/{userId}",
		UserGet,
	},
	utils.Route {
		// sign up user, add them to database as unverified, send them an email with a verifyAccount token
		"SignUp",
		"POST",
		root,
		UserSignUp,
	},
	utils.Route {
		"DeleteUser",
		"DELETE",
		root + "/{userId}",
		UserDelete,
	},
	utils.Route {
		"Login",
		"POST",
		root + "/login",
		Login,
	},
	utils.Route {
		"Logout",
		"POST",
		root + "/logout",
		Logout,
	},
	utils.Route {
		// generate reset password token and send it to provided email
		"GenerateResetPasswordToken",
		"POST",
		root + "/tokens/resetPassword",
		GenerateResetPasswordToken,
	},
	utils.Route {
		// given a token, look it up and check if it exists, and if it does, do its action
		// actions can be resetPassword or verifyAccount
		"DoActionOnce",
		"POST",
		root + "/tokens",
		OneTimeAction,
	},
}