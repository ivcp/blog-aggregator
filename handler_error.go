package main

import "net/http"

func handlerError(w http.ResponseWriter, req *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
