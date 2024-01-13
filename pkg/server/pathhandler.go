package server

import (
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": //home page
		HomeHandler(w, r)
	case "/download":
		DownloadHandler(w, r)
	case "/400":
		ErrorHandler(w, r, 400)
	case "/500":
		ErrorHandler(w, r, 500)
	default: //any invalid path will be redirected to 404 error handler
		ErrorHandler(w, r, 404)
	}
}
