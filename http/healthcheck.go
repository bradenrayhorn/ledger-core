package http

import "net/http"

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}
