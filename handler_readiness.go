package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	type resp struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, resp{Status: "ok"})
}
