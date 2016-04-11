package utils

import "net/http"

func validateAuthHeader(req *http.Request) (error, bool) {
	// check database for auth session
}
