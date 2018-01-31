package handler

import "net/http"

// handles http pings
// this is commonly used by load balancers such as
// NGinx and AWS E/ALBs
func Ping(w http.ResponseWriter, r *http.Request) {
	// Simply write a valid response header
	w.WriteHeader(200)
	return
}